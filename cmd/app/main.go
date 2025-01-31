package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"firstRest/internal/database"
	"firstRest/internal/handlers"
)

func main() {
	database.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/showTask", handlers.ShowTaskHandler).Methods("GET")
	router.HandleFunc("/addTask", handlers.AddTaskHandler).Methods("POST")
	router.HandleFunc("/updateTask/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/deleteTask/{id}", handlers.DeleteTaskHandler).Methods("DELETE")

	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
