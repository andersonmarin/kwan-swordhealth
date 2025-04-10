package notification

import (
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) NotifyTaskPerformed(t *task.Task) error {
	args := s.Called(t)
	return args.Error(0)
}
