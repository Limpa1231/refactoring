package userService

import (
	"firstRest/internal/database"
	"firstRest/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user models.User) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id uint, user models.User) (*models.User, error)
	DeleteUsers(id uint) error // Убедитесь, что этот метод есть
}

type PersonRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) AddUser(user models.User) (*models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *PersonRepository) GetAllUsers() ([]models.User, error) {
	var persons []models.User
	result := database.DB.Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}

	return persons, nil
}

func (r *PersonRepository) UpdateUser(id uint, updatedUser models.User) (*models.User, error) {
    var existingUser models.User

    
    if err := database.DB.Where("email = ? AND id <> ?", updatedUser.Email, id).First(&existingUser).Error; err == nil {
        return nil, fmt.Errorf("пользователь с email %s уже существует", updatedUser.Email)
    }

    // Находим пользователя по ID
    var person models.User
    result := database.DB.First(&person, id)
    if result.Error != nil {
        return nil, result.Error
    }

    // Обновляем поля пользователя
    person.Email = updatedUser.Email
    person.Password = updatedUser.Password // Обновляем пароль, если нужно

    // Сохраняем изменения
    result = database.DB.Save(&person)
    if result.Error != nil {
        return nil, result.Error
    }

    return &person, nil
}

func (r *PersonRepository) DeleteUsers(id uint) error {
	var persons models.User
	result := database.DB.Unscoped().First(&persons, id)
	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Delete(&persons)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
