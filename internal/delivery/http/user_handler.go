package http

import (
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
// @Param user body entity.User true "User Data"
// @Success 201 {object} entity.User
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /users/register [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.userUseCase.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

// LoginUser godoc
// @Summary Login a user
// @Description Authenticates a user and returns a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Email and Password"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /users/login [post]
func (h *UserHandler) LoginUser(c echo.Context) error {
	var credentials map[string]string
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	email := credentials["email"]
	password := credentials["password"]

	token, err := h.userUseCase.Login(c.Request().Context(), email, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
