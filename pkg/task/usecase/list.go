package usecase

import (
	"errors"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
	"github.com/andersonmarin/kwan-swordhealth/pkg/user"
	"time"
)

type ListTaskInput struct {
	UserID uint64 `json:"userId"`
}

type ListTaskOutput struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"userId"`
	Summary     string    `json:"summary"`
	PerformedAt time.Time `json:"performed_at"`
}

type ListTask struct {
	taskRepository task.Repository
	userRepository user.Repository
}

func NewListTask(taskRepository task.Repository, userRepository user.Repository) *ListTask {
	return &ListTask{taskRepository: taskRepository, userRepository: userRepository}
}

func (ct *ListTask) Execute(input *ListTaskInput) ([]*ListTaskOutput, error) {
	u, err := ct.userRepository.FindOneByID(input.UserID)
	if err != nil {
		return nil, err
	}

	var tasks []*task.Task
	if u.Role == user.RoleManager {
		tasks, err = ct.taskRepository.FindAll()
	} else if u.Role == user.RoleTechnician {
		tasks, err = ct.taskRepository.FindByUserID(u.ID)
	} else {
		return nil, errors.New("unauthorized role to list tasks")
	}
	if err != nil {
		return nil, err
	}

	output := make([]*ListTaskOutput, len(tasks))
	for i, t := range tasks {
		output[i] = &ListTaskOutput{
			ID:          t.ID,
			UserID:      t.UserID,
			Summary:     t.Summary,
			PerformedAt: t.PerformedAt,
		}
	}

	return output, nil
}
