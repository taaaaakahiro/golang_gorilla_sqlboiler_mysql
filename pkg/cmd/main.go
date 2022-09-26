package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/config"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/handler"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/persistence"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/server"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"honnef.co/go/tools/lintcmd/version"
)

const (
	exitOk    = 0
	exitError = 1
)

func main() {
	run(context.Background())
}

func run(ctx context.Context) int {
	// init logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err)
		return exitError
	}
	defer logger.Sync()
	logger = logger.With(zap.String("version", version.Version))

	// load config
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", zap.Error(err))
		return exitError
	}

	// init listener
	listener, err := net.Listen("tcp", cfg.Address())
	if err != nil {
		logger.Error("failed to listen port", zap.Int("port", cfg.Port), zap.Error(err))
		return exitError
	}
	logger.Info("server start listening", zap.Int("port", cfg.Port))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// init mysql
	logger.Info("connect to mysql ", zap.String("DSN", cfg.DB.DSN))
	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}

	mysqlDatabase, dbOpen, err := io.NewDatabase(sqlSetting)
	if err != nil {
		logger.Error("failed to create mysql db repository", zap.Error(err), zap.String("DSN", cfg.DB.DSN))
		return exitError
	}

	repositories, err := persistence.NewRepositories(mysqlDatabase, dbOpen)
	if err != nil {
		logger.Error("failed to new repositories", zap.Error(err))
		return exitError
	}
	registry := handler.NewHandler(ctx, logger, repositories, version.Version)
	httpServer := server.NewServer(ctx, registry, &server.Config{Log: logger})
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return httpServer.Serve(listener)
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
	case <-ctx.Done():
	}

	if err := httpServer.GracefulShutdown(ctx); err != nil {
		return exitError
	}

	cancel()
	if err := wg.Wait(); err != nil {
		return exitError
	}
	// //gorilla
	// r := mux.NewRouter()
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// an example API handler
	// 	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	// })
	// r.HandleFunc("/sample/{id}", sampleHandler1)

	// // サーバ設定
	// srv := &http.Server{
	// 	Handler:      r,
	// 	Addr:         "127.0.0.1:8082",
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// // 起動
	// log.Fatal(srv.ListenAndServe())
	// logger.Info("server start listening", zap.Int("port", cfg.Port))

	return exitOk
}

func sampleHandler1(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cfg, _ := config.LoadConfig(ctx)

	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}

	db, err := sql.Open("mysql", sqlSetting.SqlDsn)
	if err != nil {
		log.Fatal(err)
	}
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	user, err := models.FindUser(ctx, db, id)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}
