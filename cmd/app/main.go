package main

import (
	"log"
	"net/http"

	"firstRest/internal/database"
	"firstRest/internal/handlers"
	"firstRest/internal/taskService"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	// Создаем репозиторий, сервис и обработчики
	taskRepo := taskService.NewTaskRepository(database.DB)
	taskSvc := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskSvc)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", taskHandler.AddTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks", taskHandler.ShowTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskHandler.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", taskHandler.DeleteTaskHandler).Methods("DELETE")

	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
