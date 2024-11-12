package repository

import (
	"context"
	"tasksManagement/internal/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	GetTasks(ctx context.Context, userID string, role string) ([]*entity.Task, error)
}
