package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
)

type fetchTasksInteractor struct {
	input          dto.FetchTasksInputData
	taskRepository repository.ITaskRepository
}

func NewFetchTasksInteractor(input dto.FetchTasksInputData, taskRepository repository.ITaskRepository) *fetchTasksInteractor {
	return &fetchTasksInteractor{input: input, taskRepository: taskRepository}
}

func (i fetchTasksInteractor) Handle() (*dto.FetchTasksOutputData, error) {
	inputErrors := dto.InputErrors{}
	userID, err := vo.NewID(i.input.UserID())
	if err != nil {
		inputErrors["userID"] = err
	}

	if len(inputErrors) > 0 {
		return dto.NewFetchTasksOutputData(nil, inputErrors), nil
	}

	tasks, err := i.taskRepository.FetchTasksByUserID(*userID)
	if err != nil {
		return nil, err
	}

	return dto.NewFetchTasksOutputData(tasks, nil), nil
}
