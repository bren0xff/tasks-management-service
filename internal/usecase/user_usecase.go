package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/repository"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, email, password string) (string, error)
}

type userUseCase struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserUseCase(userRepo repository.UserRepository, jwtSecret string) UserUseCase {
	return &userUseCase{userRepo, jwtSecret}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return uc.userRepo.CreateUser(ctx, user)
}

func (uc *userUseCase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.jwtSecret))
}
