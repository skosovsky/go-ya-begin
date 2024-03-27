package repository

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"yp-examples/todo_server/internal/model"
)

func TestInMemoryRepo_AddTodo(t *testing.T) {
	r := &InMemoryRepo{}
	newTodo := model.Todo{
		Task: "test task",
		Done: false,
	}
	r.AddTodo(newTodo)

	assert.Equal(t, 1, len(r.todos))
	assert.Equal(t, newTodo.Task, r.todos[0].Task)
	assert.Equal(t, newTodo.Done, r.todos[0].Done)
}

func TestInMemoryRepo_DeleteTodo(t *testing.T) {
	todoID := uuid.NewString()

	type fields struct {
		todos []model.Todo
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{
			name: "1. Success",
			fields: fields{todos: []model.Todo{
				{
					ID:   todoID,
					Task: "Test task",
					Done: true,
				},
			}},
			args:    args{id: todoID},
			wantErr: "",
		},
		{
			name: "2. Not found error",
			fields: fields{todos: []model.Todo{
				{
					ID:   todoID,
					Task: "Test task",
					Done: true,
				},
			}},
			args:    args{id: uuid.NewString()},
			wantErr: "todo not found",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := &InMemoryRepo{
				todos: tt.fields.todos,
			}
			err := r.DeleteTodo(tt.args.id)
			if tt.wantErr != "" {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, 0, len(r.todos))
		})
	}
}

func TestInMemoryRepo_GetTodos(t *testing.T) {
	defaultTodo := model.Todo{
		ID:   uuid.NewString(),
		Task: "test task",
		Done: false,
	}
	r := &InMemoryRepo{todos: []model.Todo{defaultTodo}}
	r.GetTodos()

	assert.Equal(t, 1, len(r.todos))
	assert.Equal(t, defaultTodo.Task, r.todos[0].Task)
	assert.Equal(t, defaultTodo.Done, r.todos[0].Done)
}

func TestInMemoryRepo_UpdateTodo(t *testing.T) {
	defaultTodo := model.Todo{
		ID:   uuid.NewString(),
		Task: "test task",
		Done: false,
	}
	updatedTodo := model.Todo{
		ID:   defaultTodo.ID,
		Task: "test task",
		Done: true,
	}

	type fields struct {
		todos []model.Todo
	}
	type args struct {
		updatedTodo model.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Todo
		wantErr bool
	}{
		{
			name:    "1. Success",
			fields:  fields{todos: []model.Todo{defaultTodo}},
			args:    args{updatedTodo: updatedTodo},
			want:    updatedTodo,
			wantErr: false,
		},
		{
			name:    "2. Not found error",
			fields:  fields{todos: []model.Todo{defaultTodo}},
			args:    args{updatedTodo: model.Todo{ID: uuid.NewString()}},
			want:    model.Todo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRepo{
				todos: tt.fields.todos,
			}
			got, err := r.UpdateTodo(tt.args.updatedTodo)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
