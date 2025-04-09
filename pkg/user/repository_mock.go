package user

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) FindOneByID(id uint64) (*User, error) {
	args := r.Called(id)
	return args.Get(0).(*User), args.Error(1)
}
