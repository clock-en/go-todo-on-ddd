package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
)

type TaskRepository struct{}

func (r TaskRepository) Create(id string, title string, content string, userID string) error {
	dao := dao.TaskDao{}
	err := dao.Create(id, title, content, userID)
	return err
}
