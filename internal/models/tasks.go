package models

import "gorm.io/gorm"

type Message struct {
	ID        uint           `json:"id"`
	Task      string         `json:"task" gorm:"not null"`
	IsDone    bool           `json:"is_done" gorm:"default:false"`
	UserID    uint           `json:"user_id" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;soft_delete:0"`
}
