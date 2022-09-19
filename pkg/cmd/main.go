package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	ctx := context.Background()

	user, err := models.FindUser(ctx, db, 1)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(user.ID)
	fmt.Println(user.Name)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
