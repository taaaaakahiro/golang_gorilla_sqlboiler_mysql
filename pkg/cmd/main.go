package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/config"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
	"go.uber.org/zap"
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

	//gorilla
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	r.HandleFunc("/sample", sampleHandler1)

	// サーバ設定
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8082",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// 起動
	log.Fatal(srv.ListenAndServe())

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
	user, err := models.FindUser(ctx, db, 1)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}
