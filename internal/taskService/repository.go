package taskService

import (
	"firstRest/internal/models"
	"firstRest/internal/database"
	"gorm.io/gorm"
)

type MessageRepository interface {
	AddTaskHandler(task models.Message) (*models.Message, error)
	ShowTasksHandler() ([]models.Message, error)
	UpdateTaskHandler(id uint, updatedMessage models.Message) (*models.Message, error)
	DeleteTaskHandler(id uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) AddTaskHandler(task models.Message) (*models.Message, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *TaskRepository) ShowTasksHandler() ([]models.Message, error) {
	var messages []models.Message
	result := database.DB.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func (r *TaskRepository) UpdateTaskHandler(id uint, updatedMessage models.Message) (*models.Message, error) {
	var message models.Message
	result := database.DB.First(&message, id)
	if result.Error != nil {
		return nil, result.Error
	}

	message.Task = updatedMessage.Task
	message.IsDone = updatedMessage.IsDone

	result = database.DB.Save(&message)
	if result.Error != nil {
		return nil, result.Error
	}

	return &message, nil
}

func (r *TaskRepository) DeleteTaskHandler(id uint) error {
	var message models.Message
	result := database.DB.Unscoped().First(&message, id)
	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Delete(&message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
