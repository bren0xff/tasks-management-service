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

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*entity.User), args.Error(1)
}

type MockNotifier struct {
	mock.Mock
	wg *sync.WaitGroup
}

func (m *MockNotifier) NotifyManager(user *entity.User, task *entity.Task) {
	m.Called(user, task)
	m.wg.Done()
}

func TestTaskUseCase_CreateTask(t *testing.T) {
	mockTaskRepo := new(MockTaskRepository)
	mockUserRepo := new(MockUserRepository)
	wg := &sync.WaitGroup{}
	mockNotifier := &MockNotifier{wg: wg}
	taskUC := usecase.NewTaskUseCase(mockTaskRepo, mockUserRepo, mockNotifier)

	task := &entity.Task{
		ID:      "1",
		Summary: "Test Task",
		UserID:  "123",
	}

	user := &entity.User{
		ID:   "123",
		Role: entity.RoleTechnician,
	}

	mockUserRepo.On("GetUserByID", mock.Anything, "123").Return(user, nil)
	mockTaskRepo.On("CreateTask", mock.Anything, task).Return(nil)
	wg.Add(1)
	mockNotifier.On("NotifyManager", user, task).Return()

	err := taskUC.CreateTask(context.Background(), task)
	require.NoError(t, err)

	wg.Wait()

	mockUserRepo.AssertCalled(t, "GetUserByID", mock.Anything, "123")
	mockTaskRepo.AssertCalled(t, "CreateTask", mock.Anything, task)
	mockNotifier.AssertCalled(t, "NotifyManager", user, task)
}
