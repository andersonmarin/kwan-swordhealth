package task

type Repository interface {
	Create(task *Task) (uint64, error)
	FindAll() ([]*Task, error)
	FindByUserID(userID uint64) ([]*Task, error)
}
