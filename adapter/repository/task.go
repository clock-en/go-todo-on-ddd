package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
)

type TaskRepository struct{}

func (r TaskRepository) Create(task entity.Task) error {
	dao := dao.TaskDao{}
	err := dao.Create(task)
	return err
}
