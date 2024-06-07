package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"todo-app/database"
	"todo-app/handlers"
)

func main() {
	database.ConnectDB()

	router := chi.NewRouter()

	router.Get("/tasks", handlers.GetTasks)
	router.Post("/tasks", handlers.PostTask)
	router.Get("/tasks/{id}", handlers.GetTaskByID)
	router.Delete("/tasks/{is}", handlers.DeleteTask)

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
