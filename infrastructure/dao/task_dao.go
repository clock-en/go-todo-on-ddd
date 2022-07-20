package dao

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	_ "github.com/go-sql-driver/mysql"
)

type ITaskDao interface {
	CreateTask()
	Close()
}

type TaskDao struct {
	sqlHandler *SqlHandler
	ITaskDao
}

func (t TaskDao) CreateTask(task entity.Task) (entity.Task, error) {
	const sql = "INSERT INTO tasks (task,limitDate,status) VALUES (?, ?, ?);"
	_, err := t.sqlHandler.connect.Exec(sql, task.Title(), task.Content(), task.UserId())
	return task, err
}

func (t TaskDao) Close() {
	t.sqlHandler.connect.Close()
}

func NewTaskDao() *TaskDao {
	sqlHandler := newSqlHandler()
	return &TaskDao{sqlHandler: sqlHandler}
}
