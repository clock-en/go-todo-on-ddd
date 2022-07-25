package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type ITaskRepository interface {
	Create(entity.Task) (*entity.Task, error)
	FindById(id vo.ID) (*entity.Task, error)
}
