package handlers

import (
	"context"
	"firstRest/internal/taskService"
	"firstRest/internal/models"
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
		Message string `json:"message"`
	}

	// Парсим JSON из тела запроса
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Ошибка при разборе JSON: %v", err)})
	}

	// Добавляем задачу через сервис
	task, err := h.service.AddTask(requestBody.Message)
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

func (h *Handler) UpdateTaskHandler(c echo.Context) error {
	// Получаем ID задачи из параметров пути
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный ID"})
	}

	// Парсим JSON из тела запроса
	var updatedTask models.Message
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Ошибка при разборе JSON: %v", err)})
	}

	// Обновляем задачу через сервис
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
