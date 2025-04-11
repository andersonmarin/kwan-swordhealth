package main

import (
	"github.com/andersonmarin/kwan-swordhealth/internal/broker"
	"github.com/andersonmarin/kwan-swordhealth/internal/http"
	"github.com/andersonmarin/kwan-swordhealth/internal/mysql"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task/usecase"
	"log"
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

	if err = http.ListenAndServe(createTask, listTask); err != nil {
		log.Fatalln("HTTP listen and serve error:", err)
	}
}
