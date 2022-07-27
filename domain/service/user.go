package service

import (
	"fmt"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) userService {
	return userService{userRepository: userRepository}
}
func (s userService) ExistsName(name vo.UserName) error {
	registeredUser, _ := s.userRepository.FindUserByName(name)

	if registeredUser != nil {
		return fmt.Errorf("user name already exists")
	}
	return nil
}
func (s userService) ExistsEmail(email vo.Email) error {
	registeredUser, _ := s.userRepository.FindUserByEmail(email)

	if registeredUser != nil {
		return fmt.Errorf("email already exists")
	}
	return nil
}
