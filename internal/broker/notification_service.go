package broker

import (
	"encoding/json"
	"github.com/andersonmarin/kwan-swordhealth/pkg/notification/usecase"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
	"github.com/nats-io/nats.go"
)

type NotificationService struct {
	nc *nats.Conn
}

func NewNotificationService(nc *nats.Conn) *NotificationService {
	return &NotificationService{nc: nc}
}

func (n *NotificationService) NotifyTaskPerformed(t *task.Task) error {
	payload, err := json.Marshal(&usecase.NotifyTaskPerformedInput{
		UserID:      t.UserID,
		Summary:     t.Summary,
		PerformedAt: t.PerformedAt,
	})
	if err != nil {
		return err
	}

	return n.nc.Publish("notification.notifyTaskPerformed", payload)
}
