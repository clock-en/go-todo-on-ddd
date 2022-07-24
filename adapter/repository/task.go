package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase"
)

type ITaskDao interface {
	Create(id string, title string, content string, userID string) error
	FindById(id string) (usecase.ITaskRow, error)
}

type taskRepository struct{}

func NewTaskRepository() taskRepository {
	return taskRepository{}
}

func (r taskRepository) Create(id string, title string, content string, userID string) error {
	dao := dao.TaskDao{}
	err := dao.Create(id, title, content, userID)
	return err
}

func (q taskRepository) FindById(id string) (usecase.ITaskRow, error) {
	dao := dao.TaskDao{}
	return dao.FindById(id)
}
