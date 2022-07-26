package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type IUserRepository interface {
	CreateUser(user entity.User) (*entity.AuthUser, error)
	FindRegisteredUser(name vo.UserName, email vo.Email) (*entity.AuthUser, error)
}
