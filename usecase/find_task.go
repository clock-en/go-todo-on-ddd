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
	id, _ := vo.NewID(i.input.ID())

	task, dberror := i.taskRepository.FindTaskByID(*id)
	if dberror != nil {
		return nil, dberror
	}

	output := dto.NewFindTaskOutputData(*task)
	return output, nil
}
