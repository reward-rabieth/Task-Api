package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Priority    string
	DueDate     time.Time `gorm:"not null"`
}
