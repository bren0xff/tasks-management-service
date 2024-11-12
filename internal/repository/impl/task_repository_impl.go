package impl

import (
	"context"
	"gorm.io/gorm"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/repository"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateTask(ctx context.Context, task *entity.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	var tasks []*entity.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error) {
	var tasks []*entity.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
