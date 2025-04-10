package notification

import "github.com/andersonmarin/kwan-swordhealth/pkg/task"

type Service interface {
	NotifyTaskPerformed(t *task.Task) error
}
