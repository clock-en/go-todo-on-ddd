package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
)

type findTaskInteractor struct {
	input          dto.FindTaskInputData
	taskRepository repository.ITaskRepository
}

func NewFindTaskInteractor(input dto.FindTaskInputData, taskRepository repository.ITaskRepository) *findTaskInteractor {
	return &findTaskInteractor{input: input, taskRepository: taskRepository}
}

func (i findTaskInteractor) Handle() (*dto.FindTaskOutputData, error) {
	inputErrors := dto.InputErrors{}
	id, err := vo.NewID(i.input.ID())
	if err != nil {
		inputErrors["id"] = err
	}

	if len(inputErrors) > 0 {
		return dto.NewFindTaskOutputData(nil, inputErrors), nil
	}

	task, err := i.taskRepository.FindTaskByID(*id)
	if err != nil {
		return nil, err
	}

	return dto.NewFindTaskOutputData(task, nil), nil
}
