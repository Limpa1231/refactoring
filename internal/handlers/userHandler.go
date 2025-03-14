package handlers

import (
	"context"
	"firstRest/internal/models"
	"firstRest/internal/userService"
	"firstRest/internal/web/users"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service *userService.UserService
}

// DeleteUsersId implements users.StrictServerInterface.
func (h *userHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	panic("unimplemented")
}

// PatchUsersId implements users.StrictServerInterface.
func (h *userHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	panic("unimplemented")
}

// GetUsers implements users.StrictServerInterface.
func (h *userHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	panic("unimplemented")
}

// PostUsers implements users.StrictServerInterface.
func (h *userHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	panic("unimplemented")
}

func NewUserHandler(service *userService.UserService) *userHandler {
	return &userHandler{service: service}
}

// GetUsers возвращает всех пользователей
func (h *userHandler) ShowUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось получить пользователей"})
	}
	return c.JSON(http.StatusOK, users)
}

// PostUsers создает нового пользователя
func (h *userHandler) AddUsers(c echo.Context) error {
	var requestBody struct {
		User string `json:"user"`
		// Password int16 `json:"password"`
	}

	// Парсим JSON из тела запроса
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Ошибка при разборе JSON: %v", err)})
	}

	// Добавляем пользователя через сервис
	user, err := h.service.AddUser(requestBody.User)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось сохранить пользователя"})
	}

	// Возвращаем созданную задачу с кодом 201
	return c.JSON(http.StatusCreated, user)
}

// PatchUsersId обновляет пользователя по ID
func (h *userHandler) UpdateUsers(c echo.Context) error {
	// Получаем ID из параметров пути
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный ID"})
	}

	// Парсим JSON из тела запроса
	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Ошибка при разборе JSON: %v", err)})
	}

	// Обновляем пользователя через сервис
	user, err := h.service.UpdateUser(uint(id), updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось обновить пользователя"})
	}

	// Возвращаем обновленного пользователя
	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUsers(c echo.Context) error {
	// Получаем ID задачи из параметров пути
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный ID"})
	}

	// Удаляем задачу через сервис
	err = h.service.DeleteUsers(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось удалить пользователя"})
	}

	// Возвращаем статус 204 (No Content)
	return c.NoContent(http.StatusNoContent)
}
