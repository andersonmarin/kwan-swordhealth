package http

import (
	"errors"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task/usecase"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func ListenAndServe(createTask *usecase.CreateTask, listTask *usecase.ListTask) error {
	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		return errors.New("ADDRESS environment variable not set")
	}

	jwtKey, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		return errors.New("JWT_KEY environment variable not set")
	}

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Gzip(), middleware.CORS())

	task := e.Group("/task", echojwt.JWT([]byte(jwtKey)))
	task.POST("", NewCreateTaskHandler(createTask).Handle)
	task.GET("", NewListTaskHandler(listTask).Handle)

	return e.Start(address)
}
