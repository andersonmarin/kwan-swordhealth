package task

import "time"

const SummaryMaxLength = 2500

type Task struct {
	ID          uint64
	UserID      uint64
	Summary     string
	PerformedAt time.Time
}
