package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string         `json:"email" gorm:"unique;not null"` 
	Password  string         `json:"password" gorm:"not null"`    
	DeletedAt gorm.DeletedAt `gorm:"index;soft_delete:0"`         
}