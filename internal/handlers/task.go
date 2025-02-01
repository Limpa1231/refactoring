package handlers

import (
	"encoding/json"
	"firstRest/internal/taskService"
	"firstRest/orm"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Message string `json:"message"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при разборе JSON: %v", err), http.StatusBadRequest)
		return
	}

	task, err := h.service.AddTask(requestBody.Message)
	if err != nil {
		http.Error(w, "Не удалось сохранить задачу", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) ShowTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		http.Error(w, "Не удалось получить задачи", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var updatedTask orm.Message
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedTask)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при разборе JSON: %v", err), http.StatusBadRequest)
		return
	}

	task, err := h.service.UpdateTask(uint(id), updatedTask)
	if err != nil {
		http.Error(w, "Не удалось обновить задачу", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(uint(id))
	if err != nil {
		http.Error(w, "Не удалось удалить задачу", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
