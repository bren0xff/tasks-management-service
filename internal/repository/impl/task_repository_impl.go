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

func (r *taskRepository) GetTasks(ctx context.Context, userID int64, role string) ([]*entity.Task, error) {
	var tasks []*entity.Task
	if role == "manager" {
		err := r.db.Find(&tasks).Error
		return tasks, err
	}
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
