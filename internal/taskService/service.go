package taskService

import (
	"firstRest/orm"
)

type TaskService struct {
	repo MessageRepository
}

func NewTaskService(repo MessageRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(task string) (*orm.Message, error) {
	message := orm.Message{
		Task:   task,
		IsDone: false,
	}
	return s.repo.AddTaskHandler(message)
}

func (s *TaskService) GetAllTasks() ([]orm.Message, error) {
	return s.repo.ShowTasksHandler()
}

func (s *TaskService) UpdateTask(id uint, updatedTask orm.Message) (*orm.Message, error) {
	return s.repo.UpdateTaskHandler(id, updatedTask)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskHandler(id)
}
