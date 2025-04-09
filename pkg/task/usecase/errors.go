package usecase

import "errors"

var (
	ErrUserNotFound               = errors.New("user not found")
	ErrUserNotAllowedToCreateTask = errors.New("only technicians can create tasks")
	ErrPerformedAtInFuture        = errors.New("performedAt cannot be in the future")
	ErrSummaryTooLong             = errors.New("summary exceeds 2500 characters")
)
