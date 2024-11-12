package usecase

import (
	"context"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/notifier"
	"tasksManagement/internal/repository"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *entity.Task, user entity.User) error
	GetTasks(ctx context.Context, userID string, role entity.Role) ([]*entity.Task, error)
}

type taskUseCase struct {
	taskRepo repository.TaskRepository
	notifier notifier.Notifier
}

func NewTaskUseCase(tr repository.TaskRepository, n notifier.Notifier) TaskUseCase {
	return &taskUseCase{tr, n}
}

func (uc *taskUseCase) CreateTask(ctx context.Context, task *entity.Task, user entity.User) error {
	if user.Role == entity.RoleTechnician {
		go uc.notifier.NotifyManager(task)
	}
	return uc.taskRepo.CreateTask(ctx, task)
}

func (uc *taskUseCase) GetTasks(ctx context.Context, userID string, role entity.Role) ([]*entity.Task, error) {
	if role == entity.RoleManager {
		return uc.taskRepo.GetAllTasks(ctx)
	}
	return uc.taskRepo.GetTasksByUserID(ctx, userID)
}
