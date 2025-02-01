package taskService

import (
	"firstRest/internal/database"
	"firstRest/orm"
	"gorm.io/gorm"
)

type MessageRepository interface {
	AddTaskHandler(task orm.Message) (*orm.Message, error)
	ShowTasksHandler() ([]orm.Message, error)
	UpdateTaskHandler(id uint, updatedMessage orm.Message) (*orm.Message, error)
	DeleteTaskHandler(id uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) AddTaskHandler(task orm.Message) (*orm.Message, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *TaskRepository) ShowTasksHandler() ([]orm.Message, error) {
	var messages []orm.Message
	result := database.DB.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func (r *TaskRepository) UpdateTaskHandler(id uint, updatedMessage orm.Message) (*orm.Message, error) {
	var message orm.Message
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
	var message orm.Message
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
