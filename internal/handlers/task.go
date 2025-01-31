package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"firstRest/internal/database"
	"firstRest/orm"
	"github.com/gorilla/mux"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var requestBody struct {
			Message string `json:"message"`
		}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при разборе JSON: %v", err), http.StatusBadRequest)
			return
		}

		task := requestBody.Message
		fmt.Fprintln(w, "Задача успешно сохранена:", task)

		message := orm.Message{
			Task:   task,
			IsDone: false,
		}

		result := database.DB.Create(&message)
		if result.Error != nil {
			fmt.Println("Ошибка при сохранении записи в базу данных:", result.Error)
			http.Error(w, "Не удалось сохранить задачу", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func ShowTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var messages []orm.Message
		database.DB.Find(&messages)
		json.NewEncoder(w).Encode(messages)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		params := mux.Vars(r)
		id := params["id"]

		var message orm.Message
		result := database.DB.First(&message, id)

		if result.Error != nil {
			http.Error(w, "Запись не найдена", http.StatusNotFound)
			return
		}

		var updatedMessage orm.Message
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updatedMessage)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при разборе JSON: %v", err), http.StatusBadRequest)
			return
		}

		message.Task = updatedMessage.Task
		message.IsDone = updatedMessage.IsDone

		result = database.DB.Save(&message)
		if result.Error != nil {
			http.Error(w, "Не удалось обновить запись", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		params := mux.Vars(r)
		id := params["id"]

		log.Printf("Полученный ID: %s\n", id)

		var message orm.Message
		result := database.DB.Unscoped().First(&message, id)

		if result.Error != nil {
			log.Printf("Запись не найдена: %v\n", result.Error)
			http.Error(w, "Запись не найдена", http.StatusNotFound)
			return
		}

		log.Printf("Найденная запись: %+v\n", message)

		result = database.DB.Delete(&message)
		if result.Error != nil {
			log.Printf("Не удалось удалить запись: %v\n", result.Error)
			http.Error(w, "Не удалось удалить запись", http.StatusInternalServerError)
			return
		}

		log.Printf("Запись успешно удалена\n")
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
