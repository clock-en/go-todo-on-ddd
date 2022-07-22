package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/google/uuid"
)

type createTaskInput struct {
	title   vo.TaskTitle
	content vo.TaskContent
	userID  vo.ID
}

func NewCreateTaskInput(title string, content string, userID string) createTaskInput {
	t, _ := vo.NewTaskTitle(title)
	c, _ := vo.NewTaskContent(content)
	u, _ := vo.NewID(userID)
	return createTaskInput{title: *t, content: *c, userID: *u}
}

type CreateTaskOutput struct {
	task entity.Task
}

func (o CreateTaskOutput) Task() entity.Task {
	return o.task
}

type createTaskInteractor struct {
	input          createTaskInput
	taskRepository repository.TaskRepository
}

func NewCreateTaskUsecase(input createTaskInput) *createTaskInteractor {
	return &createTaskInteractor{input: input, taskRepository: repository.TaskRepository{}}
}

func (i createTaskInteractor) Handler() (*CreateTaskOutput, error) {
	generatedID, _ := uuid.NewRandom()
	id, err := vo.NewID(generatedID.String())
	if err != nil {
		return nil, err
	}
	task := entity.NewTask(*id, i.input.title, i.input.content, i.input.userID)
	dberror := i.taskRepository.Create(*task)
	if dberror != nil {
		return nil, dberror
	}
	return &CreateTaskOutput{task: *task}, nil
}
