package entity

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type User struct {
	id    vo.ID
	name  vo.UserName
	email vo.Email
}

type Users []User

func NewUser(id vo.ID, name vo.UserName, email vo.Email) *User {
	return &User{id: id, name: name, email: email}
}

func (u User) ID() vo.ID {
	return u.id
}

func (u User) Name() vo.UserName {
	return u.name
}

func (u User) Email() vo.Email {
	return u.email
}
