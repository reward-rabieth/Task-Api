package Task

import (
	"context"
	"github.con/reward-rabieth/Task-Api/core/components/Task/models"
	"gorm.io/gorm"
)

type Repo interface {
	CreateTask(ctx context.Context, task *models.Task) error
	GetTask(ctx context.Context, id uint) (*models.Task, error)
	UpdateTask(ctx context.Context, id uint, task *models.Task) error
	DeleteTask(ctx context.Context, id uint) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (r *repo, err error) {
	r = &repo{db: db}
	return
}

func (s *repo) CreateTask(ctx context.Context, task *models.Task) error {
	return s.db.Create(task).Error
}

func (s *repo) GetTask(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	err := s.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *repo) UpdateTask(ctx context.Context, id uint, task *models.Task) error {
	// Find the task by its ID
	result := s.db.First(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}

	// Update the task fields
	result = s.db.Model(&models.Task{}).Where("id = ?", id).Updates(task)
	return result.Error
}

func (s *repo) DeleteTask(ctx context.Context, id uint) error {
	// Find and delete the task by its ID
	result := s.db.Delete(&models.Task{}, id)
	return result.Error
}
