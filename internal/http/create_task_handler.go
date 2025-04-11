package http

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/task/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateTaskHandler struct {
	Handler
	createTask *usecase.CreateTask
}

func NewCreateTaskHandler(createTask *usecase.CreateTask) *CreateTaskHandler {
	return &CreateTaskHandler{createTask: createTask}
}

func (h *CreateTaskHandler) Handle(c echo.Context) error {
	userID, err := h.currentUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	var input usecase.CreateTaskInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	input.UserID = userID

	output, err := h.createTask.Execute(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}
