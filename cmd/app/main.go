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
	router.HandleFunc("/addTask", taskHandler.AddTaskHandler).Methods("POST")
	router.HandleFunc("/showTasks", taskHandler.ShowTasksHandler).Methods("GET")
	router.HandleFunc("/updateTask/{id}", taskHandler.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/deleteTask/{id}", taskHandler.DeleteTaskHandler).Methods("DELETE")

	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
