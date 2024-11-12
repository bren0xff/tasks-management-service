package usecase

import (
	"context"
	"errors"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/notifier"
	"tasksManagement/internal/repository"
	"time"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	GetTasks(ctx context.Context, userID string, role string) ([]*entity.Task, error)
}

type taskUseCase struct {
	taskRepo repository.TaskRepository
	notifier notifier.Notifier
}

func NewTaskUseCase(tr repository.TaskRepository, n notifier.Notifier) TaskUseCase {
	return &taskUseCase{tr, n}
}

func (uc *taskUseCase) CreateTask(ctx context.Context, task *entity.Task) error {
	if len(task.Summary) > 2500 {
		return errors.New("summary exceeds 2500 characters")
	}
	task.Date = time.Now().Format("2006-01-02")
	err := uc.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return err
	}
	go uc.notifier.NotifyManager(task)
	return nil
}

func (uc *taskUseCase) GetTasks(ctx context.Context, userID string, role string) ([]*entity.Task, error) {
	return uc.taskRepo.GetTasks(ctx, userID, role)
}
