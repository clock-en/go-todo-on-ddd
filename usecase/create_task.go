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
	generatedID, _ := uuid.NewRandom()
	id, _ := vo.NewID(generatedID.String())
	title, _ := vo.NewTaskTitle(i.input.Title())
	content, _ := vo.NewTaskContent(i.input.Content())
	userID, _ := vo.NewID(i.input.UserID())
	newTask := entity.NewTask(*id, *title, *content, *userID)

	task, dberror := i.taskRepository.CreateTask(*newTask)
	if dberror != nil {
		return nil, dberror
	}

	output := dto.NewCreateTaskOutputData(*task)
	return output, nil
}
