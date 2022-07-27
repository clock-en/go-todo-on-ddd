package usecase

import (
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
	name, email, password, inputErrors := i.validateInputs()
	if len(inputErrors) > 0 {
		return dto.NewRegisterUserOutputData(nil, inputErrors), nil
	}

	generatedID, _ := uuid.NewRandom()
	id, err := vo.NewID(generatedID.String())
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(*id, *name, *email, *password)

	authUser, err := i.userRepository.CreateUser(*user)
	if err != nil {
		return nil, err
	}

	return dto.NewRegisterUserOutputData(authUser, nil), nil
}
func (i registerUserInteractor) validateInputs() (*vo.UserName, *vo.Email, *vo.Password, dto.InputErrors) {
	inputErrors := dto.InputErrors{}
	name, err := vo.NewUserName(i.input.Name())
	if err != nil {
		inputErrors["name"] = err
	}
	email, err := vo.NewEmail(i.input.Email())
	if err != nil {
		inputErrors["email"] = err
	}
	password, err := vo.NewPassword(i.input.Password())
	if err != nil {
		inputErrors["password"] = err
	}

	service := service.NewUserService(i.userRepository)
	if err := service.ExistsName(*name); err != nil {
		inputErrors["name"] = err
	}
	if err := service.ExistsEmail(*email); err != nil {
		inputErrors["email"] = err
	}

	return name, email, password, inputErrors
}
