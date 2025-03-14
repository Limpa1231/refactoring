package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task      string         `json:"task"`
	IsDone    bool           `json:"is_done"`
	DeletedAt gorm.DeletedAt `gorm:"index;soft_delete:0"`
}
