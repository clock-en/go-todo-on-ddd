package dao

import (
	_ "github.com/go-sql-driver/mysql"
)

type taskRow struct {
	id      string
	title   string
	content string
	userID  string
}

func (r taskRow) ID() string {
	return r.id
}
func (r taskRow) Title() string {
	return r.title
}
func (r taskRow) Content() string {
	return r.content
}
func (r taskRow) UserID() string {
	return r.userID
}

type TaskDao struct {
	handler *sqlHandler
}

func (t TaskDao) Create(id string, title string, content string, userID string) error {
	handler := newSqlHandler()
	defer handler.connect.Close()
	const sql = "INSERT INTO tasks (id, title, content, user_id) VALUES (?, ?, ?, ?);"
	_, err := handler.connect.Exec(sql, id, title, content, userID)
	return err
}

func (t TaskDao) FindById(taskID string) (*taskRow, error) {
	handler := newSqlHandler()
	defer handler.connect.Close()
	const sql = "SELECT * FROM tasks WHERE id=?;"
	row := handler.connect.QueryRow(sql, taskID)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var id, title, content, userID string
	row.Scan(&id, &title, &content, &userID)
	return &taskRow{id, title, content, userID}, nil
}
