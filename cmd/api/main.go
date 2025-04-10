package main

import (
	"github.com/andersonmarin/kwan-swordhealth/internal/broker"
	"github.com/andersonmarin/kwan-swordhealth/internal/mysql"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	db, err := mysql.OpenConnection()
	if err != nil {
		log.Fatalln("MySQL connection error:", err)
	}
	defer db.Close()

	nc, err := broker.OpenNatsConnection()
	if err != nil {
		log.Fatalln("Nats connection error:", err)
	}
	defer nc.Close()

	taskRepository := mysql.NewTaskRepository(db)
	userRepository := mysql.NewUserRepository(db)
	notificationService := broker.NewNotificationService(nc)

	createTask := usecase.NewCreateTask(taskRepository, userRepository, notificationService)
	listTask := usecase.NewListTask(taskRepository, userRepository)

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Gzip(), middleware.CORS())

	e.POST("/task", func(c echo.Context) error {
		var input usecase.CreateTaskInput
		if err := c.Bind(&input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		input.UserID = 2

		output, err := createTask.Execute(&input)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, output)
	})

	e.GET("/task", func(c echo.Context) error {
		output, err := listTask.Execute(&usecase.ListTaskInput{
			UserID: 2,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return echo.NewHTTPError(http.StatusOK, output)
	})

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
