package mysql

import (
	"database/sql"
	"github.com/andersonmarin/kwan-swordhealth/pkg/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindOneByID(id uint64) (*user.User, error) {
	rows, err := ur.db.Query(`
	SELECT id, username, password, role 
	FROM users 
	WHERE id = ? 
	LIMIT 1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var u user.User
	if err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role); err != nil {
		return nil, err
	}

	return &u, nil
}
