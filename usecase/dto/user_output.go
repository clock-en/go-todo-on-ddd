package dto

import "github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"

type RegisterUserOutputData struct {
	user userData
}

func NewRegisterUserOutputData(user entity.AuthUser) *RegisterUserOutputData {
	return &RegisterUserOutputData{
		user: createUserData(user),
	}
}
func (o RegisterUserOutputData) Task() userData {
	return o.user
}

type userData struct {
	id    string
	name  string
	email string
}

func (u userData) ID() string {
	return u.id
}
func (u userData) Name() string {
	return u.name
}
func (u userData) Email() string {
	return u.email
}
func createUserData(user entity.AuthUser) userData {
	return userData{
		id:    user.ID().Value(),
		name:  user.Name().Value(),
		email: user.Email().Value(),
	}
}
