package handlers

import (
	"bytes"
	"encoding/json"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodoHandler(t *testing.T) {
	todo := models.Todo{
		ID:       "1",
		Titulo:   "Meu Primeiro ToDo",
		Compelto: false,
	}

	body, _ := json.Marshal(todo)

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseWriter simulado
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler retornou código de status incorreto: obteve %v, esperava %v", status, http.StatusCreated)
	}

	var result models.Todo
	if err := json.Unmarshal(rr.Body.Bytes(), &result); err != nil {
		t.Fatalf("Falha ao analisar JSON: %v", err)
	}

	if result.ID != todo.ID || result.Titulo != todo.Titulo || result.Compelto != todo.Compelto {
		t.Errorf("handler retornou dados inesperados: obteve %v, esperava %v", result, todo)
	}
}

func TestGetTodosHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTodos)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código de status incorreto: obteve %v, esperava %v", status, http.StatusOK)
	}

	var todos []models.Todo
	if err := json.Unmarshal(rr.Body.Bytes(), &todos); err != nil {
		t.Fatalf("Falha ao analisar JSON: %v", err)
	}
}

func TestCreateTodoInvalidData(t *testing.T) {
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTodo)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler retornou código de status incorreto: obteve %v, esperava %v", status, http.StatusBadRequest)
	}
}

func TestTodoRepositoryCreate(t *testing.T) {
	repo := repository.NewTodoRepository()
	todo := models.Todo{
		ID:       "2",
		Titulo:   "Meu Segundo ToDo",
		Compelto: false,
	}
	repo.Create(todo)

	todos := repo.GetAll()
	if len(todos) != 1 {
		t.Errorf("esperava 1 ToDo na lista, obteve %d", len(todos))
	}
	if todos[0].ID != todo.ID {
		t.Errorf("esperava ToDo com ID igual a %s, mas obteve %s", todo.ID, todos[0].ID)
	}
}

func TestTodoRepositoryGetAll(t *testing.T) {
	repo := repository.NewTodoRepository()
	repo.Create(models.Todo{ID: "3", Titulo: "README", Compelto: true})

	todos := repo.GetAll()
	if len(todos) != 1 {
		t.Errorf("esperava 1 ToDo no repositório, mas obteve %d", len(todos))
	}
}

func TestCreateMultipleTodos(t *testing.T) {
	repo := repository.NewTodoRepository()

	todo1 := models.Todo{ID: "1", Titulo: "Todo 1", Compelto: false}
	todo2 := models.Todo{ID: "2", Titulo: "Todo 2", Compelto: true}

	repo.Create(todo1)
	repo.Create(todo2)

	todos := repo.GetAll()
	if len(todos) != 2 {
		t.Errorf("esperava 2 Todos, mas obteve %d", len(todos))
	}
}
