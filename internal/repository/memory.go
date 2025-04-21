package repository

import "go-rest-api/internal/models"

type TodoRepository struct {
	todos map[string]models.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: make(map[string]models.Todo),
	}
}

func (r *TodoRepository) Create(todo models.Todo) {
	r.todos[todo.ID] = todo
}

func (r *TodoRepository) GetAll() []models.Todo {
	var todos []models.Todo
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos
}
