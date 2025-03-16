package taskService

import (
	"firstRest/internal/database"
	"firstRest/internal/models"
	"fmt"
)

type TaskService struct {
	repo MessageRepository
}

func NewTaskService(repo MessageRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(title string, isDone bool, userID uint) (*models.Message, error) {
	// Проверяем, существует ли пользователь
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return nil, fmt.Errorf("пользователь с ID %d не найден", userID)
	}

	// Создаем задачу
	task := models.Message{
		Task:   title,
		IsDone: isDone,
		UserID: userID,
	}

	result = database.DB.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}
func (s *TaskService) GetAllTasks() ([]models.Message, error) {
	return s.repo.ShowTasksHandler()
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]models.Message, error) {
	var tasks []models.Message
	result := database.DB.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(id uint, updatedTask models.Message) (*models.Message, error) {
	var task models.Message
	result := database.DB.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Обновляем поля задачи
	task.Task = updatedTask.Task
	task.IsDone = updatedTask.IsDone
	task.UserID = updatedTask.UserID

	// Сохраняем изменения
	result = database.DB.Save(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskHandler(id)
}
