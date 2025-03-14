package main

import (
	"firstRest/internal/database"
	"firstRest/internal/handlers"
	"firstRest/internal/taskService"
	"firstRest/internal/userService"
	"firstRest/internal/web/tasks"
	"firstRest/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Создаем репозиторий, сервис и обработчики для задач
	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewTaskService(tasksRepo)
	taskHandler := handlers.NewTaskHandler(tasksService)
	// Создаем репозиторий, сервис и обработчики для пользователей
	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewUserService(usersRepo)
	userHandler := handlers.NewUserHandler(usersService)
	// Инициализация Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Логирование запросов
	e.Use(middleware.Recover()) // Восстановление после паник

	// Регистрация хендлеров через oapi-codegen
	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	// Регистрация хендлеров для пользователей через oapi-codegen
	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	e.POST("/api/tasks", taskHandler.AddTaskHandler)
	e.GET("/api/tasks", taskHandler.ShowTasksHandler)
	e.PUT("/api/tasks/:id", taskHandler.UpdateTaskHandler)
	e.DELETE("/api/tasks/:id", taskHandler.DeleteTaskHandler)

	// Регистрация хендлеров для пользователей
	e.GET("/api/users", userHandler.ShowUsers)
	e.POST("/api/users", userHandler.AddUsers)
	e.DELETE("/api/users/:id", userHandler.DeleteUsers)
	e.PUT("/api/users/:id", userHandler.UpdateUsers)

	// Запуск сервера
	log.Println("Server started at localhost:8080")
	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
