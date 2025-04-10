package usecase

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/notification"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
	"github.com/andersonmarin/kwan-swordhealth/pkg/user"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestCreateTask_Execute(t *testing.T) {
	t.Run("should create a new task", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)
		notificationService := new(notification.ServiceMock)

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

		notificationService.On("NotifyTaskPerformed", &task.Task{
			ID:          taskID,
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		}).Return(nil)

		output, err := NewCreateTask(taskRepository, userRepository, notificationService).Execute(&CreateTaskInput{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		})

		require.NoError(t, err)
		require.Equal(t, &CreateTaskOutput{ID: taskID}, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should fail if user does not exist", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)
		notificationService := new(notification.ServiceMock)

		var (
			userID      = uint64(7)
			summary     = "Some task"
			performedAt = time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC)
		)

		userRepository.On("FindOneByID", uint64(7)).Return((*user.User)(nil), nil)

		output, err := NewCreateTask(taskRepository, userRepository, notificationService).Execute(&CreateTaskInput{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		})

		require.ErrorIs(t, err, ErrUserNotFound)
		require.Nil(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should fail if user role is not allowed", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)
		notificationService := new(notification.ServiceMock)

		var (
			userID      = uint64(7)
			summary     = "Some task"
			performedAt = time.Date(9025, time.April, 8, 0, 0, 0, 0, time.UTC)
		)

		userRepository.On("FindOneByID", uint64(7)).Return(&user.User{
			ID:   userID,
			Role: user.RoleManager,
		}, nil)

		output, err := NewCreateTask(taskRepository, userRepository, notificationService).Execute(&CreateTaskInput{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		})

		require.ErrorIs(t, err, ErrUserNotAllowedToCreateTask)
		require.Nil(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should fail if summary is empty", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)
		notificationService := new(notification.ServiceMock)

		var (
			userID      = uint64(7)
			summary     = ""
			performedAt = time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC)
		)

		userRepository.On("FindOneByID", uint64(7)).Return(&user.User{
			ID:   userID,
			Role: user.RoleTechnician,
		}, nil)

		output, err := NewCreateTask(taskRepository, userRepository, notificationService).Execute(&CreateTaskInput{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		})

		require.ErrorIs(t, err, ErrSummaryEmpty)
		require.Nil(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should fail if summary exceeds max length", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)
		notificationService := new(notification.ServiceMock)

		var (
			userID      = uint64(7)
			summary     = strings.Repeat("A", 2501)
			performedAt = time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC)
		)

		userRepository.On("FindOneByID", uint64(7)).Return(&user.User{
			ID:   userID,
			Role: user.RoleTechnician,
		}, nil)

		output, err := NewCreateTask(taskRepository, userRepository, notificationService).Execute(&CreateTaskInput{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		})

		require.ErrorIs(t, err, ErrSummaryTooLong)
		require.Nil(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})

	t.Run("should fail if performedAt is in the future", func(t *testing.T) {
		taskRepository := new(task.RepositoryMock)
		userRepository := new(user.RepositoryMock)
		notificationService := new(notification.ServiceMock)

		var (
			userID      = uint64(7)
			summary     = "Some task"
			performedAt = time.Date(9025, time.April, 8, 0, 0, 0, 0, time.UTC)
		)

		userRepository.On("FindOneByID", uint64(7)).Return(&user.User{
			ID:   userID,
			Role: user.RoleTechnician,
		}, nil)

		output, err := NewCreateTask(taskRepository, userRepository, notificationService).Execute(&CreateTaskInput{
			UserID:      userID,
			Summary:     summary,
			PerformedAt: performedAt,
		})

		require.ErrorIs(t, err, ErrPerformedAtInFuture)
		require.Nil(t, output)

		taskRepository.AssertExpectations(t)
		userRepository.AssertExpectations(t)
	})
}
