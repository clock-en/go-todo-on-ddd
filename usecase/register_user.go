package usecase

import (
	"fmt"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/service"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
	"github.com/google/uuid"
)

type registerUserInteractor struct {
	input          dto.RegisterUserInputData
	userRepository repository.IUserRepository
}

func NewRegisterUserInteractor(
	input dto.RegisterUserInputData,
	userRepository repository.IUserRepository,
) *registerUserInteractor {
	return &registerUserInteractor{input: input, userRepository: userRepository}
}
func (i registerUserInteractor) Handle() (*dto.RegisterUserOutputData, error) {
	generatedID, _ := uuid.NewRandom()
	id, _ := vo.NewID(generatedID.String())
	name, _ := vo.NewUserName(i.input.Name())
	email, _ := vo.NewEmail(i.input.Email())
	password, _ := vo.NewPassword(i.input.Password())
	user := entity.NewUser(*id, *name, *email, *password)

	service := service.NewUserService(i.userRepository)
	if service.Exists(*user) {
		return nil, fmt.Errorf("user already exists")
	}

	authUser, dberror := i.userRepository.CreateUser(*user)
	if dberror != nil {
		return nil, dberror
	}

	output := dto.NewRegisterUserOutputData(*authUser)
	return output, nil
}
