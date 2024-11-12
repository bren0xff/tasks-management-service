package http

import (
	"github.com/google/uuid"
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
// @Param input body CreateTaskRequest true "Task Input Data"
// @Success 201 {object} entity.Task
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /tasks [post]
// @Security BearerAuth
func (h *TaskHandler) CreateTask(c echo.Context) error {
	var input CreateTaskRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request data",
		})
	}
	user := c.Get("user").(*entity.User)

	task := entity.Task{
		ID:      uuid.New().String(),
		Summary: input.Summary,
		Date:    input.Date,
		UserID:  user.ID,
	}

	if err := h.taskUseCase.CreateTask(c.Request().Context(), &task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
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
// @Failure 401 {object} map[string]interface{} "Unauthorized"
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

type CreateTaskRequest struct {
	Summary string `json:"summary" example:"Fix server issue"`
	Date    string `json:"date" example:"2024-11-12"`
}
