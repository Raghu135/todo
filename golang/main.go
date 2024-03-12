package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=todo_user1 dbname=todo_database sslmode=require")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// ROUTEs
	router.HandleFunc("/todos", GetTodosHandler(db)).Methods("GET")
	router.HandleFunc("/todos", CreateTodoHandler(db)).Methods("POST")
	router.HandleFunc("/todos/{id:[0-9]+}", EditTodoHandler(db)).Methods("PUT")
	router.HandleFunc("/todos/{id:[0-9]+}", DeleteTodoHandler(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
