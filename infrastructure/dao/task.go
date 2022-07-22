package dao

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	_ "github.com/go-sql-driver/mysql"
)

type TaskDao struct {
	handler *sqlHandler
}

func (t TaskDao) Create(task entity.Task) error {
	handler := newSqlHandler()
	defer handler.connect.Close()
	const sql = "INSERT INTO tasks (id, title, content, user_id) VALUES (?, ?, ?, ?);"
	_, err := handler.connect.Exec(sql, task.ID().Value(), task.Title().Value(), task.Content().Value(), task.UserID().Value())
	return err
}
