package nats

import (
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
	//TODO implement me
	panic("implement me")
}
