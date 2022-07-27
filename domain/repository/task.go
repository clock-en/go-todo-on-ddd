package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type ITaskRepository interface {
	CreateTask(entity.Task) error
	FindTaskByID(id vo.ID) (*entity.Task, error)
	FetchTasksByUserID(userID vo.ID) (entity.Tasks, error)
}
