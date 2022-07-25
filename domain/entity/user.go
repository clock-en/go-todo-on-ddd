package entity

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type User struct {
	id       vo.ID
	name     vo.UserName
	email    vo.Email
	password vo.Password
}

func NewUser(id vo.ID, name vo.UserName, email vo.Email, password vo.Password) *User {
	return &User{id: id, name: name, email: email, password: password}
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

func (u User) Password() vo.Password {
	return u.password
}

type AuthUser struct {
	id    vo.ID
	name  vo.UserName
	email vo.Email
}

func NewAuthUser(id vo.ID, name vo.UserName, email vo.Email) *User {
	return &User{id: id, name: name, email: email}
}

func (u AuthUser) ID() vo.ID {
	return u.id
}

func (u AuthUser) Name() vo.UserName {
	return u.name
}

func (u AuthUser) Email() vo.Email {
	return u.email
}
