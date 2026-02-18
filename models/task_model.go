package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Des       string     `json:"des"`
	Completed bool       `json:"completed"`
	DueDate   *time.Time `json:"due_date"`

	UserID uint `json:"user_id"`
	User   User `json:"-" gorm:"foreignKey:UserID"`
}
