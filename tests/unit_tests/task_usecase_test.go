package unit_tests

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"sync"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/usecase"
	"testing"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(ctx context.Context, task *entity.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*entity.Task), args.Error(1)
}

type MockNotifier struct {
	mock.Mock
	wg *sync.WaitGroup
}

func (m *MockNotifier) NotifyManager(task *entity.Task) {
	m.Called(task)
	m.wg.Done()
}

func TestTaskUseCase_CreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	wg := &sync.WaitGroup{}
	mockNotifier := &MockNotifier{wg: wg}
	taskUC := usecase.NewTaskUseCase(mockRepo, mockNotifier)

	task := &entity.Task{
		ID:      "1",
		Summary: "Test Task",
	}

	user := &entity.User{
		ID:   "123",
		Role: entity.RoleTechnician,
	}

	mockRepo.On("CreateTask", mock.Anything, task).Return(nil)
	wg.Add(1)
	mockNotifier.On("NotifyManager", task).Return()

	err := taskUC.CreateTask(context.Background(), task, *user)
	require.NoError(t, err)
	wg.Wait()

	mockRepo.AssertCalled(t, "CreateTask", mock.Anything, task)
	mockNotifier.AssertCalled(t, "NotifyManager", task)
}
