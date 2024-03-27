package repository

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"yp-examples/todo_server/internal/model"
)

type InMemoryRepo struct {
	mu    sync.Mutex
	todos []model.Todo
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{}
}

func (r *InMemoryRepo) GetTodos() []model.Todo {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.todos
}

func (r *InMemoryRepo) AddTodo(todo model.Todo) model.Todo {
	r.mu.Lock()
	defer r.mu.Unlock()
	todo.ID = uuid.NewString()
	r.todos = append(r.todos, todo)
	return todo
}

func (r *InMemoryRepo) UpdateTodo(updatedTodo model.Todo) (model.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, todo := range r.todos {
		if todo.ID == updatedTodo.ID {
			r.todos[i].Done = updatedTodo.Done
			return updatedTodo, nil
		}
	}
	return model.Todo{}, fmt.Errorf("todo not found")
}

func (r *InMemoryRepo) DeleteTodo(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, todo := range r.todos {
		if todo.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}
