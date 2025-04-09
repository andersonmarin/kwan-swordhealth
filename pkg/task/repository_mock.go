package task

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Create(task *Task) (*Task, error) {
	args := r.Called(task)
	return args.Get(0).(*Task), args.Error(1)
}

func (r *RepositoryMock) FindAll() ([]*Task, error) {
	args := r.Called()
	return args.Get(0).([]*Task), args.Error(1)
}

func (r *RepositoryMock) FindByUserID(userID uint64) ([]*Task, error) {
	args := r.Called(userID)
	return args.Get(0).([]*Task), args.Error(1)
}
