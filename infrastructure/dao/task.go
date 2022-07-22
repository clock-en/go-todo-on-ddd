package dao

import (
	_ "github.com/go-sql-driver/mysql"
)

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
