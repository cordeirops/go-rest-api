package handlers

import (
	"encoding/json"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
)

var repo = repository.NewTodoRepository()

func SetupTodoRoutes(r *mux.Router) {
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todos", CreateTodo).Methods("POST")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := repo.GetAll()
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo.ID == "" || todo.Titulo == "" {
		http.Error(w, "Campos obrigatórios estão faltando", http.StatusBadRequest)
		return
	}

	repo.Create(todo)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
