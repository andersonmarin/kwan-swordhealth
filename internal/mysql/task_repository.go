package mysql

import (
	"database/sql"
	"github.com/andersonmarin/kwan-swordhealth/pkg/task"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) Create(t *task.Task) (uint64, error) {
	result, err := tr.db.Exec(`
	INSERT INTO tasks(user_id, summary, performed_at) 
	VALUES (?, ?, ?)
	`, t.UserID, t.Summary, t.PerformedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (tr *TaskRepository) FindAll() ([]*task.Task, error) {
	rows, err := tr.db.Query(`
	SELECT id, user_id, summary, performed_at 
	FROM tasks 
	ORDER BY performed_at
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*task.Task, 0)
	for rows.Next() {
		var t task.Task
		if err = rows.Scan(&t.ID, &t.UserID, &t.Summary, &t.PerformedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}

func (tr *TaskRepository) FindByUserID(userID uint64) ([]*task.Task, error) {
	rows, err := tr.db.Query(`
	SELECT id, user_id, summary, performed_at 
	FROM tasks
	WHERE user_id = ?
	ORDER BY performed_at
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*task.Task, 0)
	for rows.Next() {
		var t task.Task
		if err = rows.Scan(&t.ID, &t.UserID, &t.Summary, &t.PerformedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}
