package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
	"net/http"
)

var repo = repository.NewTodoRepository()

func SetupTodoRoutes(r *mux.Router) {
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todos", CreateTodo).Methods("POST")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := repo.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	repo.Create(todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
