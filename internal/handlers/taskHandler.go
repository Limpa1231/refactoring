package handlers

import (
	"context"
	"firstRest/internal/database"
	"firstRest/internal/models"
	"firstRest/internal/taskService"
	"firstRest/internal/web/tasks"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *taskService.TaskService
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	panic("unimplemented")
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	panic("unimplemented")
}

func NewTaskHandler(service *taskService.TaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddTaskHandler(c echo.Context) error {
	var requestBody struct {
		Task   string `json:"Task"`
		IsDone bool   `json:"Is_done"`
		UserID uint   `json:"User_Id"`
	}

	// Парсим JSON из тела запроса
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Ошибка при разборе JSON: %v", err)})
	}

	// Создаем задачу через сервис
	task, err := h.service.AddTask(requestBody.Task, requestBody.IsDone, requestBody.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось сохранить задачу"})
	}

	// Возвращаем созданную задачу с кодом 201
	return c.JSON(http.StatusCreated, task)
}

func (h *Handler) ShowTasksHandler(c echo.Context) error {
	// Получаем все задачи через сервис
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось получить задачи"})
	}

	// Возвращаем список задач
	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTasksByUserID(c echo.Context) error {
	// Получаем user_id из параметров пути
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный user_id"})
	}

	// Получаем задачи через сервис
	tasks, err := h.service.GetTasksByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось получить задачи"})
	}

	// Возвращаем список задач
	return c.JSON(http.StatusOK, tasks)
}

// GetTasks implements tasks.StrictServerInterface.

func (h *Handler) UpdateTaskHandler(c echo.Context) error {
	// Получаем ID задачи из параметров пути
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный ID"})
	}

	// Парсим JSON из тела запроса
	var requestBody struct {
		Task   string `json:"task"`
		IsDone bool   `json:"is_done"`
		UserID uint   `json:"user_id"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Ошибка при разборе JSON: %v", err)})
	}

	// Проверяем, существует ли пользователь с указанным UserID
	var user models.User
	if err := database.DB.First(&user, requestBody.UserID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Пользователь с ID %d не найден", requestBody.UserID)})
	}

	// Обновляем задачу через сервис
	updatedTask := models.Message{
		Task:   requestBody.Task,
		IsDone: requestBody.IsDone,
		UserID: requestBody.UserID,
	}
	task, err := h.service.UpdateTask(uint(id), updatedTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось обновить задачу"})
	}

	// Возвращаем обновленную задачу
	return c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTaskHandler(c echo.Context) error {
	// Получаем ID задачи из параметров пути
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный ID"})
	}

	// Удаляем задачу через сервис
	err = h.service.DeleteTask(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось удалить задачу"})
	}

	// Возвращаем статус 204 (No Content)
	return c.NoContent(http.StatusNoContent)
}
