package task

type Repository interface {
	Create(task *Task) (*Task, error)
	FindAll() ([]*Task, error)
	FindByUserID(userID uint64) ([]*Task, error)
}
