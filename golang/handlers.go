// handler

package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

// Todo struct
type Todo struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Data      string    `json:"data" db:"data"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

// GETS
func GetTodosHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func CreateTodoHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		json.NewDecoder(r.Body).Decode(&todo)

		// Insert
		_, err := db.Exec("INSERT INTO todos (title, data, expires_at) VALUES ($1, $2, $3)", todo.Title, todo.Data, todo.ExpiresAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todo)
	}
}

// MOdify
func EditTodoHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// REMOVE
func DeleteTodoHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
