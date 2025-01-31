package taskService

import (
	"firstRest/internal/database"
	"firstRest/orm"
)

func AddTask(task string) (*orm.Message, error) {
	message := orm.Message{
		Task:   task,
		IsDone: false,
	}

	result := database.DB.Create(&message)
	if result.Error != nil {
		return nil, result.Error
	}

	return &message, nil
}

func GetTasks() ([]orm.Message, error) {
	var messages []orm.Message
	result := database.DB.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func UpdateTask(id uint, updatedMessage orm.Message) (*orm.Message, error) {
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

func DeleteTask(id uint) error {
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
