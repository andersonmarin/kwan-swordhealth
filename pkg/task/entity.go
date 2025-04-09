package task

import "time"

type Task struct {
	ID          uint64
	UserID      uint64
	Summary     string
	PerformedAt time.Time
}
