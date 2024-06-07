package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{ //nolint:gochecknoglobals // example
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postman",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func postTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	tasks[task.ID] = task

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "задачи с таким ID не существует", http.StatusBadRequest)

		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "задачи с таким ID не существует. Не могу удалить", http.StatusBadRequest)

		return
	}

	delete(tasks, task.ID)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := chi.NewRouter()

	r.Get("/tasks", getTasks)
	r.Post("/tasks", postTask)
	r.Get("/tasks/{id}", getTask)
	r.Delete("/tasks/{id}", deleteTask)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())

		return
	}
}
