package usecase

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/notification"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
	"github.com/andersonmarin/kwan-swordhealth/pkg/user"
	"time"
)

type CreateTaskInput struct {
	UserID      uint64    `json:"userId"`
	Summary     string    `json:"summary"`
	PerformedAt time.Time `json:"performedAt"`
}

type CreateTaskOutput struct {
	ID uint64 `json:"id"`
}

type CreateTask struct {
	taskRepository      task.Repository
	userRepository      user.Repository
	notificationService notification.Service
}

func NewCreateTask(taskRepository task.Repository, userRepository user.Repository, notificationService notification.Service) *CreateTask {
	return &CreateTask{taskRepository: taskRepository, userRepository: userRepository, notificationService: notificationService}
}

func (ct *CreateTask) Execute(input *CreateTaskInput) (*CreateTaskOutput, error) {
	u, err := ct.userRepository.FindOneByID(input.UserID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, task.ErrUserNotFound
	}

	if u.Role != user.RoleTechnician {
		return nil, task.ErrUserNotAllowedToCreateTask
	}

	t := task.Task{
		UserID:      input.UserID,
		Summary:     input.Summary,
		PerformedAt: input.PerformedAt,
	}

	if err = t.Validate(); err != nil {
		return nil, err
	}

	taskID, err := ct.taskRepository.Create(&t)
	if err != nil {
		return nil, err
	}
	t.ID = taskID

	if err = ct.notificationService.NotifyTaskPerformed(&t); err != nil {
		return nil, err
	}

	return &CreateTaskOutput{ID: t.ID}, nil
}
