package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func New() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/example?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
