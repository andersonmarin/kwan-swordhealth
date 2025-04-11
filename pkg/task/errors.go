package task

import "errors"

var (
	ErrPerformedAtInFuture         = errors.New("performedAt cannot be in the future")
	ErrSummaryTooLong              = errors.New("summary exceeds 2500 characters")
	ErrSummaryEmpty                = errors.New("summary is empty")
	ErrPerformedAtEmpty            = errors.New("performedAt is empty")
	ErrUnauthorizedRoleToListTasks = errors.New("unauthorized role to list tasks")
)
