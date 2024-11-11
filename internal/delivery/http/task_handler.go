package http

import (
	"net/http"
	"tasksManagement/internal/entity"
	"tasksManagement/internal/usecase"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(e *echo.Echo, tu usecase.TaskUseCase, jwtSecret string) {
	handler := &TaskHandler{tu}
	authMiddleware := AuthMiddleware(jwtSecret)

	e.POST("/tasks", handler.CreateTask, authMiddleware)
	e.GET("/tasks", handler.GetTasks, authMiddleware)
}

// CreateTask godoc
// @Summary Create a new task
// @Description Allows a technician to create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body entity.Task true "Task Data"
// @Success 201 {object} entity.Task
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Router /tasks [post]
// @Security BearerAuth
func (h *TaskHandler) CreateTask(c echo.Context) error {
	user := c.Get("user").(*entity.User)
	var task entity.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserID = user.ID
	if err := h.taskUseCase.CreateTask(c.Request().Context(), &task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, task)
}

// GetTasks godoc
// @Summary Get tasks
// @Description Retrieves tasks based on user role
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} entity.Task
// @Failure 401 {object} echo.HTTPError
// @Router /tasks [get]
// @Security BearerAuth
func (h *TaskHandler) GetTasks(c echo.Context) error {
	user := c.Get("user").(*entity.User)
	tasks, err := h.taskUseCase.GetTasks(c.Request().Context(), user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}
