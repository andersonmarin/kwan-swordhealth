package http

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/task/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ListTaskHandler struct {
	listTask *usecase.ListTask
}

func NewListTaskHandler(listTask *usecase.ListTask) *ListTaskHandler {
	return &ListTaskHandler{listTask: listTask}
}

func (h *ListTaskHandler) Handle(c echo.Context) error {
	output, err := h.listTask.Execute(&usecase.ListTaskInput{
		UserID: 2,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return echo.NewHTTPError(http.StatusOK, output)
}
