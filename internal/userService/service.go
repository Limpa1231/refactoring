package userService

import (
	"firstRest/internal/models"
	"fmt"
)

type UserService struct {
	usersRepo UserRepository
}

func (s *UserService) DeleteUser(u uint) any {
	panic("unimplemented")
}

func NewUserService(usersRepo UserRepository) *UserService {
	return &UserService{usersRepo: usersRepo}
}

func (s *UserService) AddUser(user string) (*models.User, error) {
	message := models.User{
		Email:    user,
		Password: "root",
	}
	return s.usersRepo.AddUser(message)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.usersRepo.GetAllUsers()
}

func (s *UserService) UpdateUser(id uint, updatedUser models.User) (*models.User, error) {
	return s.usersRepo.UpdateUser(id, updatedUser)
}

func (s *UserService) DeleteUsers(id uint) error {
	// Логика удаления пользователя
	if err := s.usersRepo.DeleteUsers(id); err != nil {
		return fmt.Errorf("не удалось удалить пользователя: %w", err)
	}
	return nil
}
