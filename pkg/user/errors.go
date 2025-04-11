package user

import "errors"

var (
	ErrUserNotFound               = errors.New("user not found")
	ErrUserNotAllowedToCreateTask = errors.New("only technicians can create tasks")
)
