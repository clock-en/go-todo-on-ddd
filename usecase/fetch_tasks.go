package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
)

type fetchTasksInteractor struct {
	input          dto.FetchTasksInputData
	taskRepository ITaskRepository
}

func NewFetchTasksInteractor(input dto.FetchTasksInputData, taskRepository ITaskRepository) *fetchTasksInteractor {
	return &fetchTasksInteractor{input: input, taskRepository: taskRepository}
}

func (i fetchTasksInteractor) Handle() (*dto.FetchTasksOutputData, error) {
	userID, _ := vo.NewID(i.input.UserID())

	tasks, dberror := i.taskRepository.FetchTasksByUserID(*userID)
	if dberror != nil {
		return nil, dberror
	}

	output := dto.NewFetchTasksOutputData(tasks)
	return output, nil
}
