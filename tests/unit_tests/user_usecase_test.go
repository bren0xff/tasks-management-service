package unit_tests

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/usecase"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestUserUseCase_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	jwtSecret := os.Getenv("JWT_SECRET")
	userUC := usecase.NewUserUseCase(mockRepo, jwtSecret)

	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		Role:     "technician",
	}

	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	err := userUC.CreateUser(context.Background(), user)
	require.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

func TestUserUseCase_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	jwtSecret := os.Getenv("JWT_SECRET")
	userUC := usecase.NewUserUseCase(mockRepo, jwtSecret)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	require.NoError(t, err)

	mockUser := &entity.User{
		ID:       "1",
		Email:    "test@example.com",
		Password: string(hashedPassword),
		Role:     "manager",
	}

	mockRepo.On("GetUserByEmail", mock.Anything, "test@example.com").Return(mockUser, nil)

	token, err := userUC.Login(context.Background(), "test@example.com", "password")
	require.NoError(t, err)
	require.NotEmpty(t, token)
}
