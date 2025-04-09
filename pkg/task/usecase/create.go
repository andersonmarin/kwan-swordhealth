package usecase

import (
	"errors"
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
	taskRepository task.Repository
	userRepository user.Repository
}

func NewCreateTask(taskRepository task.Repository, userRepository user.Repository) *CreateTask {
	return &CreateTask{taskRepository: taskRepository, userRepository: userRepository}
}

func (ct *CreateTask) Execute(input *CreateTaskInput) (*CreateTaskOutput, error) {
	u, err := ct.userRepository.FindOneByID(input.UserID)
	if err != nil {
		return nil, err
	}

	if u.Role != user.RoleTechnician {
		return nil, errors.New("only technicians can create tasks")
	}

	if input.PerformedAt.After(time.Now()) {
		return nil, errors.New("performedAt cannot be in the future")
	}

	t, err := ct.taskRepository.Create(&task.Task{
		UserID:      input.UserID,
		Summary:     input.Summary,
		PerformedAt: input.PerformedAt,
	})
	if err != nil {
		return nil, err
	}

	return &CreateTaskOutput{ID: t.ID}, nil
}
