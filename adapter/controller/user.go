package controller

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
)

type userController struct {
	dao repository.IUserDao
}

func NewUserController(dao repository.IUserDao) userController {
	return userController{dao: dao}
}
func (c userController) RegisterUser(name string, email string, password string) (*dto.RegisterUserOutputData, error) {
	input := dto.NewRegisterUserInputData(name, email, password)
	userRepository := repository.NewUserRepository(c.dao)
	interactor := usecase.NewRegisterUserInteractor(input, userRepository)
	return interactor.Handle()
}

type userViewModel struct {
	id    string
	name  string
	email string
}

func (v userViewModel) ID() string {
	return v.id
}
func (v userViewModel) Name() string {
	return v.name
}
func (v userViewModel) Email() string {
	return v.email
}
