package usecase

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
	"github.com/andersonmarin/kwan-swordhealth/pkg/user"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListTask_Execute(t *testing.T) {
	t.Run("should list only technician tasks", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)

		userID := uint64(7)

		taskRepository.On("FindByUserID", userID).Return([]*task.Task{
			{
				ID:          42,
				UserID:      userID,
				Summary:     "Some task",
				PerformedAt: time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC),
			},
		}, nil)

		userRepository.On("FindOneByID", userID).Return(&user.User{
			ID:   userID,
			Role: user.RoleTechnician,
		}, nil)

		output, err := NewListTask(taskRepository, userRepository).Execute(&ListTaskInput{
			UserID: userID,
		})
		require.NoError(t, err)
		require.NotEmpty(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should list only all tasks", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)

		userID := uint64(7)

		taskRepository.On("FindAll").Return([]*task.Task{
			{
				ID:          42,
				UserID:      userID,
				Summary:     "Some task",
				PerformedAt: time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC),
			},
		}, nil)

		userRepository.On("FindOneByID", userID).Return(&user.User{
			ID:   userID,
			Role: user.RoleManager,
		}, nil)

		output, err := NewListTask(taskRepository, userRepository).Execute(&ListTaskInput{
			UserID: userID,
		})
		require.NoError(t, err)
		require.NotEmpty(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should fail if user does not exist", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)

		userID := uint64(7)

		userRepository.On("FindOneByID", userID).Return((*user.User)(nil), nil)

		output, err := NewListTask(taskRepository, userRepository).Execute(&ListTaskInput{
			UserID: userID,
		})
		require.ErrorIs(t, err, user.ErrUserNotFound)
		require.Nil(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should fail if role is unauthorized", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)

		userID := uint64(7)

		userRepository.On("FindOneByID", userID).Return(&user.User{
			ID:   userID,
			Role: "trainee",
		}, nil)

		output, err := NewListTask(taskRepository, userRepository).Execute(&ListTaskInput{
			UserID: userID,
		})
		require.ErrorIs(t, err, task.ErrUnauthorizedRoleToListTasks)
		require.Nil(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})
}
