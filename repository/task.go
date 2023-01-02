package repository

import (
	"context"
	"e-learning/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTaskByUserId(ctx context.Context, id int) (entity.Question, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) GetTaskByUserId(ctx context.Context, id int) (entity.Question, error) {
	return entity.Question{}, nil
}
