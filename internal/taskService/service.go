package taskService

import "firstRest/internal/models"

type TaskService struct {
	repo MessageRepository
}

func NewTaskService(repo MessageRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(task string) (*models.Message, error) {
	message := models.Message{
		Task:   task,
		IsDone: false,
	}
	return s.repo.AddTaskHandler(message)
}

func (s *TaskService) GetAllTasks() ([]models.Message, error) {
	return s.repo.ShowTasksHandler()
}

func (s *TaskService) UpdateTask(id uint, updatedTask models.Message) (*models.Message, error) {
	return s.repo.UpdateTaskHandler(id, updatedTask)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskHandler(id)
}
