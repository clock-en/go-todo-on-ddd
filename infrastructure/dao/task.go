package dao

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	_ "github.com/go-sql-driver/mysql"
)

type TaskDao struct {
	handler *sqlHandler
}

func (t TaskDao) CreateTask(task entity.Task) (entity.Task, error) {
	const sql = "INSERT INTO tasks (id, title, content, user_id) VALUES (?, ?, ?, ?);"
	_, err := t.handler.connect.Exec(sql, task.ID().Value(), task.Title().Value(), task.Content().Value(), task.UserID().Value())
	return task, err
}

func (t TaskDao) Close() {
	t.handler.connect.Close()
}

func NewTaskDao() *TaskDao {
	sqlHandler := newSqlHandler()
	return &TaskDao{handler: sqlHandler}
}
