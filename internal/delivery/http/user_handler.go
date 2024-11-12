package http

import (
	"github.com/google/uuid"
	"net/http"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(e *echo.Echo, userUseCase usecase.UserUseCase) {
	handler := &UserHandler{userUseCase}
	e.POST("/users/register", handler.RegisterUser)
	e.POST("/users/login", handler.LoginUser)
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Creates a new user with a hashed password
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterUserRequest true "User Data (excluding ID)"
// @Success 201 {object} entity.User
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users/register [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request data",
		})
	}

	role := entity.Role(input.Role)
	if !role.IsValid() {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid role",
		})
	}

	newUser := entity.User{
		ID:       uuid.New().String(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     role,
	}

	err := h.userUseCase.CreateUser(c.Request().Context(), &newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	newUser.Password = ""
	return c.JSON(http.StatusCreated, newUser)
}

// LoginUser godoc
// @Summary Login a user
// @Description Authenticates a user and returns a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Email and Password"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /users/login [post]
func (h *UserHandler) LoginUser(c echo.Context) error {
	var credentials LoginRequest
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request data",
		})
	}

	token, err := h.userUseCase.Login(c.Request().Context(), credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Invalid credentials",
		})
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}

type RegisterUserRequest struct {
	Name     string `json:"name" example:"Fulano"`
	Email    string `json:"email" example:"fulano.sobre@example.com"`
	Password string `json:"password" example:"securepassword"`
	Role     string `json:"role" example:"technician"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password" example:"securepassword"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
