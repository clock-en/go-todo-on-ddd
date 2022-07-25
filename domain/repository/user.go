package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type IUserRepository interface {
	CreateUser(user entity.User) (*entity.AuthUser, error)
	FindByEmail(email vo.Email) (*entity.AuthUser, error)
}
