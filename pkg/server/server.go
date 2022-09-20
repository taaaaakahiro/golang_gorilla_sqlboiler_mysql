package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/config"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
)

type Server struct {
	Database *sql.DB
}

func NewServer(port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}
	db, err := io.NewDatabase(cfg.DB.Dsn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	user, err := models.FindUser(ctx, db.Database, 1)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Fprintf(w, user.Name)
	fmt.Println(user.Name)

	// fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
