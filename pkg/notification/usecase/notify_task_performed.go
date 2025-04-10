package usecase

import (
	"log"
	"time"
)

type NotifyTaskPerformedInput struct {
	UserID      uint64
	Summary     string
	PerformedAt time.Time
}

type NotifyTaskPerformed struct{}

func NewNotifyTaskPerformed() *NotifyTaskPerformed {
	return &NotifyTaskPerformed{}
}

func (ntp *NotifyTaskPerformed) Execute(input *NotifyTaskPerformedInput) error {
	log.Printf("[NOTIFIER] The tech #%d performed the task \"%s\" on date %s\n", input.UserID, input.Summary, input.PerformedAt.Format(time.DateTime))
	return nil
}
