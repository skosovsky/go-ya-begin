package handler

import (
	"encoding/json"
	"net/http"

	"yp-examples/todo_server/internal/model"
	"yp-examples/todo_server/internal/repository"
)

type TodosHandler struct {
	repo repository.Repository
}

func NewTodosHandler(repo repository.Repository) TodosHandler {
	return TodosHandler{
		repo: repo,
	}
}

func (h TodosHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(h.repo.GetTodos())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h TodosHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo := h.repo.AddTodo(todo)
	err := json.NewEncoder(w).Encode(newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h TodosHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo, err := h.repo.UpdateTodo(todo)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h TodosHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.repo.DeleteTodo(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
