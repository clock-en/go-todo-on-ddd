package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/google/uuid"
)

type createTaskInputData struct {
	title   string
	content string
	userID  string
}

func NewCreateTaskInputData(title string, content string, userID string) createTaskInputData {
	return createTaskInputData{title: title, content: content, userID: userID}
}

type createTaskOutputData struct {
	id      string
	title   string
	content string
	userID  string
}

func (o createTaskOutputData) ID() string {
	return o.id
}
func (o createTaskOutputData) Title() string {
	return o.title
}
func (o createTaskOutputData) Content() string {
	return o.content
}
func (o createTaskOutputData) UserID() string {
	return o.userID
}

type createTaskInteractor struct {
	input          createTaskInputData
	taskRepository ITaskRepository
}

func NewCreateTaskUsecase(input createTaskInputData, taskRepository ITaskRepository) *createTaskInteractor {
	return &createTaskInteractor{input: input, taskRepository: taskRepository}
}

func (i createTaskInteractor) Handle() (*createTaskOutputData, error) {
	generatedID, _ := uuid.NewRandom()
	id, _ := vo.NewID(generatedID.String())
	title, _ := vo.NewTaskTitle(i.input.title)
	content, _ := vo.NewTaskContent(i.input.content)
	userID, _ := vo.NewID(i.input.userID)
	task := entity.NewTask(*id, *title, *content, *userID)

	dberror := i.taskRepository.Create(
		task.ID().Value(),
		task.Title().Value(),
		task.Content().Value(),
		task.UserID().Value(),
	)
	if dberror != nil {
		return nil, dberror
	}

	output := &createTaskOutputData{
		id:      task.ID().Value(),
		title:   task.Title().Value(),
		content: task.Content().Value(),
		userID:  task.UserID().Value(),
	}
	return output, nil
}
