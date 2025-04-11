package http

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/task/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateTaskHandler struct {
	createTask *usecase.CreateTask
}

func NewCreateTaskHandler(createTask *usecase.CreateTask) *CreateTaskHandler {
	return &CreateTaskHandler{createTask: createTask}
}

func (cth *CreateTaskHandler) Handle(c echo.Context) error {
	var input usecase.CreateTaskInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	input.UserID = 2

	output, err := cth.createTask.Execute(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, output)
}
