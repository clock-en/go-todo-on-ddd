package service

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/repository"
)

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) userService {
	return userService{userRepository: userRepository}
}
func (s userService) Exists(user entity.User) bool {
	registeredUser, err := s.userRepository.FindByEmail(user.Email())
	if err != nil || registeredUser != nil {
		return true
	}
	return false
}
