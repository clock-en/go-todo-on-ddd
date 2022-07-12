package entity

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

// User represents entity
type User struct {
	id       vo.Id
	name     vo.UserName
	email    vo.Email
	password vo.Password
}

type Users []User

func NewUser(id vo.Id, name vo.UserName, email vo.Email, password vo.Password) *User {
	return &User{id: id, name: name, email: email, password: password}
}

func (u User) Id() vo.Id {
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
