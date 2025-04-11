package task

import "time"

const SummaryMaxLength = 2500

type Task struct {
	ID          uint64
	UserID      uint64
	Summary     string
	PerformedAt time.Time
}

func (t *Task) Validate() error {
	if len(t.Summary) == 0 {
		return ErrSummaryEmpty
	}

	if len(t.Summary) > SummaryMaxLength {
		return ErrSummaryTooLong
	}

	if t.PerformedAt.IsZero() {
		return ErrPerformedAtEmpty
	}

	if t.PerformedAt.After(time.Now()) {
		return ErrPerformedAtInFuture
	}

	return nil
}
