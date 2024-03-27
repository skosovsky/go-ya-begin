package todo_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	handler2 "yp-examples/todo_server/internal/handler"
	"yp-examples/todo_server/internal/model"
	"yp-examples/todo_server/mocks"
)

func TestGetTodos(t *testing.T) {
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	r := mocks.NewMockRepository(ctrl)
	r.EXPECT().GetTodos().Times(1).Return([]model.Todo{})
	h := handler2.NewTodosHandler(r)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetTodos)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "[]"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestAddTodo(t *testing.T) {
	todo := model.Todo{
		Task: "Test task",
		Done: false,
	}
	jsonTodo, _ := json.Marshal(todo)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonTodo))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	r := mocks.NewMockRepository(ctrl)
	r.EXPECT().AddTodo(todo).Times(1)
	h := handler2.NewTodosHandler(r)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.AddTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateTodo(t *testing.T) {
	todo := model.Todo{
		ID:   uuid.NewString(),
		Task: "Test task",
		Done: true,
	}

	ctrl := gomock.NewController(t)
	r := mocks.NewMockRepository(ctrl)
	r.EXPECT().UpdateTodo(todo).Times(1).Return(todo, nil)
	h := handler2.NewTodosHandler(r)

	td, _ := json.Marshal(todo)
	updateReq, _ := http.NewRequest("PUT", "/todos", bytes.NewBuffer(td))
	updateRr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdateTodo)
	handler.ServeHTTP(updateRr, updateReq)

	if status := updateRr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var updatedTodo model.Todo
	err := json.NewDecoder(updateRr.Body).Decode(&updatedTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, todo.Done, updatedTodo.Done, "handler returned unexpected body: got %+v want %+v", todo.Done, updatedTodo.Done)
}

func TestDeleteTodo(t *testing.T) {
	todoID := uuid.NewString()

	ctrl := gomock.NewController(t)
	r := mocks.NewMockRepository(ctrl)
	r.EXPECT().DeleteTodo(todoID).Times(1).Return(nil)
	h := handler2.NewTodosHandler(r)

	deleteReq, _ := http.NewRequest("DELETE", fmt.Sprintf("/todos?id=%s", todoID), nil)
	deleteRr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DeleteTodo)
	handler.ServeHTTP(deleteRr, deleteReq)

	if status := deleteRr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}
