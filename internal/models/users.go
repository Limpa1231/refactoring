package models

import "gorm.io/gorm"

type User struct {
	ID        uint           `json:"id"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Messages  []Message      `json:"messages" gorm:"foreignKey:UserID"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;soft_delete:0"`
}
