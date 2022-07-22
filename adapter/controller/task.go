package controller

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase"
)

type taskCreateViewModel struct {
	id      string
	title   string
	content string
	userID  string
}

func (v taskCreateViewModel) ID() string {
	return v.id
}
func (v taskCreateViewModel) Title() string {
	return v.title
}
func (v taskCreateViewModel) Content() string {
	return v.content
}
func (v taskCreateViewModel) UserID() string {
	return v.userID
}

type TaskController struct{}

func (c TaskController) Create(title string, content string, userID string) (*taskCreateViewModel, error) {
	input := usecase.NewCreateTaskInputData(title, content, userID)
	taskRepository := repository.TaskRepository{}
	interactor := usecase.NewCreateTaskUsecase(input, taskRepository)
	output, err := interactor.Handle()
	taskViewModel := &taskCreateViewModel{
		id:      output.ID(),
		title:   output.Title(),
		content: output.Content(),
		userID:  output.UserID(),
	}
	return taskViewModel, err
}
