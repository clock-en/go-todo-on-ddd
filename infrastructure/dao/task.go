package dao

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	_ "github.com/go-sql-driver/mysql"
)

type TaskDao struct {
	handler *sqlHandler
}

func (t TaskDao) CreateTask(task entity.Task) (entity.Task, error) {
	const sql = "INSERT INTO tasks (task,limitDate,status) VALUES (?, ?, ?);"
	_, err := t.handler.connect.Exec(sql, task.Title(), task.Content(), task.UserID())
	return task, err
}

func (t TaskDao) Close() {
	t.handler.connect.Close()
}

func NewTaskDao() *TaskDao {
	sqlHandler := newSqlHandler()
	return &TaskDao{handler: sqlHandler}
}
