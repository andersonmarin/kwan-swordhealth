package user

type Role string

const (
	RoleTechnician Role = "technician"
	RoleManager    Role = "manager"
)

type User struct {
	ID       uint64
	Username string
	Password string
	Role     Role
}

func (u *User) CanCreateTask() bool {
	return u.Role == RoleTechnician
}
