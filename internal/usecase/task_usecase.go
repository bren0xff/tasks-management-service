package usecase

import (
	"context"
	"errors"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/notifier"
	"tasksManagement/internal/repository"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	GetTasks(ctx context.Context, userID string, role entity.Role) ([]*entity.Task, error)
}

type taskUseCase struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
	notifier notifier.Notifier
}

func NewTaskUseCase(tr repository.TaskRepository, ur repository.UserRepository, n notifier.Notifier) TaskUseCase {
	return &taskUseCase{tr, ur, n}
}

func (uc *taskUseCase) CreateTask(ctx context.Context, task *entity.Task) error {

	user, err := uc.userRepo.GetUserByID(ctx, task.UserID)
	if err != nil {
		return errors.New("failed to fetch user")
	}
	if user.Role == entity.RoleTechnician {
		go uc.notifier.NotifyManager(user, task)
	}
	return uc.taskRepo.CreateTask(ctx, task)
}

func (uc *taskUseCase) GetTasks(ctx context.Context, userID string, role entity.Role) ([]*entity.Task, error) {
	if role == entity.RoleManager {
		return uc.taskRepo.GetAllTasks(ctx)
	}
	return uc.taskRepo.GetTasksByUserID(ctx, userID)
}
