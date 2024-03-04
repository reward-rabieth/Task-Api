package Reminder

import (
	"context"
	"github.con/reward-rabieth/Task-Api/core/components/Reminder/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo interface {
	CreateReminder(ctx context.Context, reminder *models.Reminder) (*models.Reminder, error)
	GetReminder(ctx context.Context, id uint) (*models.Reminder, error)
	UpdateReminder(ctx context.Context, id uint, reminder *models.Reminder) error
	DeleteReminder(ctx context.Context, id uint) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (r *repo, err error) {
	r = &repo{db: db}
	return
}

func (s *repo) CreateReminder(ctx context.Context, reminder *models.Reminder) (*models.Reminder, error) {
	result := s.db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(reminder)
	return reminder, result.Error
}

func (s *repo) GetReminder(ctx context.Context, id uint) (*models.Reminder, error) {
	var reminder models.Reminder
	err := s.db.First(&reminder, id).Error
	if err != nil {
		return nil, err
	}
	return &reminder, nil
}

func (s *repo) UpdateReminder(ctx context.Context, id uint, reminder *models.Reminder) error {
	result := s.db.First(&models.Reminder{}, id)
	if result.Error != nil {
		return result.Error
	}

	result = s.db.Model(&models.Reminder{}).Where("id = ?", id).Updates(reminder)
	return result.Error
}

func (s *repo) DeleteReminder(ctx context.Context, id uint) error {
	result := s.db.Delete(&models.Reminder{}, id)
	return result.Error
}
