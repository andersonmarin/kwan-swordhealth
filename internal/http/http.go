package http

import (
	"errors"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func ListenAndServe(createTask *usecase.CreateTask, listTask *usecase.ListTask) error {
	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		return errors.New("ADDRESS environment variable not set")
	}

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Gzip(), middleware.CORS())

	e.POST("/task", NewCreateTaskHandler(createTask).Handle)
	e.GET("/task", NewListTaskHandler(listTask).Handle)

	return e.Start(address)
}
