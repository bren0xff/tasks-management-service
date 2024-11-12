package repository

import (
	"context"
	"tasksManagement/internal/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	GetAllTasks(ctx context.Context) ([]*entity.Task, error)
	GetTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error)
}
