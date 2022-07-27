package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
	"github.com/google/uuid"
)

type createTaskInteractor struct {
	input          dto.CreateTaskInputData
	taskRepository repository.ITaskRepository
}

func NewCreateTaskInteractor(input dto.CreateTaskInputData, taskRepository repository.ITaskRepository) *createTaskInteractor {
	return &createTaskInteractor{input: input, taskRepository: taskRepository}
}

func (i createTaskInteractor) Handle() (*dto.CreateTaskOutputData, error) {
	title, content, userID, inputErrors := i.validateInputs()
	if len(inputErrors) > 0 {
		return dto.NewCreateTaskOutputData(nil, inputErrors), nil
	}

	generatedID, _ := uuid.NewRandom()
	id, err := vo.NewID(generatedID.String())
	if err != nil {
		return nil, err
	}
	task := entity.NewTask(*id, *title, *content, *userID)

	dbErr := i.taskRepository.CreateTask(*task)
	if dbErr != nil {
		return nil, err
	}

	return dto.NewCreateTaskOutputData(task, nil), nil
}
func (i createTaskInteractor) validateInputs() (*vo.TaskTitle, *vo.TaskContent, *vo.ID, dto.InputErrors) {
	inputErrors := dto.InputErrors{}
	title, err := vo.NewTaskTitle(i.input.Title())
	if err != nil {
		inputErrors["title"] = err
	}
	content, err := vo.NewTaskContent(i.input.Content())
	if err != nil {
		inputErrors["content"] = err
	}
	userID, err := vo.NewID(i.input.UserID())
	if err != nil {
		inputErrors["userID"] = err
	}

	return title, content, userID, inputErrors
}
