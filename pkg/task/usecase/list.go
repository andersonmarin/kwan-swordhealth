package usecase

import (
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
	PerformedAt time.Time `json:"performedAt"`
}

type ListTask struct {
	taskRepository task.Repository
	userRepository user.Repository
}

func NewListTask(taskRepository task.Repository, userRepository user.Repository) *ListTask {
	return &ListTask{taskRepository: taskRepository, userRepository: userRepository}
}

func (lt *ListTask) Execute(input *ListTaskInput) ([]*ListTaskOutput, error) {
	u, err := lt.userRepository.FindOneByID(input.UserID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, user.ErrUserNotFound
	}

	tasks, err := lt.findTasks(u)
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

func (lt *ListTask) findTasks(u *user.User) ([]*task.Task, error) {
	switch u.Role {
	case user.RoleManager:
		return lt.taskRepository.FindAll()
	case user.RoleTechnician:
		return lt.taskRepository.FindByUserID(u.ID)
	default:
		return nil, task.ErrUnauthorizedRoleToListTasks
	}
}
