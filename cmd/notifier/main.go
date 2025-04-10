package main

import (
	"encoding/json"
	"github.com/andersonmarin/kwan-swordhealth/internal/broker"
	"github.com/andersonmarin/kwan-swordhealth/pkg/notification/usecase"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	nc, err := broker.OpenNatsConnection()
	if err != nil {
		log.Fatalln("Nats connection error:", err)
	}
	defer nc.Close()

	notifyTaskPerformed := usecase.NewNotifyTaskPerformed()

	_, err = nc.Subscribe("notification.notifyTaskPerformed", func(msg *nats.Msg) {
		var input usecase.NotifyTaskPerformedInput
		if err := json.Unmarshal(msg.Data, &input); err != nil {
			log.Println("[notifyTaskPerformed] Unmarshal error:", err)
		}

		if err := notifyTaskPerformed.Execute(&input); err != nil {
			log.Println("[notifyTaskPerformed] Execute error:", err)
		}
	})
	if err != nil {
		panic(err)
	}

	select {}
}
