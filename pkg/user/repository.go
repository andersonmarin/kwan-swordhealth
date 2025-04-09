package user

type Repository interface {
	FindOneByID(id uint64) (*User, error)
}
