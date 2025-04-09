package usecase

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
	"github.com/andersonmarin/kwan-swordhealth/pkg/user"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateTask_Execute(t *testing.T) {
	t.Run("should create a new task", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)

		var (
			userID      = uint64(7)
			taskID      = uint64(42)
			summary     = "Some task"
			performedAt = time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC)
		)

		taskRepository.On("Create", &task.Task{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		}).Return(&task.Task{
			ID:          taskID,
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		}, nil)

		userRepository.On("FindOneByID", uint64(7)).Return(&user.User{
			ID:   userID,
			Role: user.RoleTechnician,
		}, nil)

		output, err := NewCreateTask(taskRepository, userRepository).Execute(&CreateTaskInput{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		})

		require.NoError(t, err)
		require.Equal(t, &CreateTaskOutput{ID: taskID}, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})
}
