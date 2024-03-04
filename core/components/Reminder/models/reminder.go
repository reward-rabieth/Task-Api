package models

import (
	"gorm.io/gorm"
	"time"
)

type Reminder struct {
	gorm.Model
	ReminderTime time.Time `gorm:"not null"`
	ContactName  string    `gorm:"not null"`
	TaskID       uint      `gorm:"index"`
}
