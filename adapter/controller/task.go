package controller

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase"
)

type taskViewModel struct {
	id      string
	title   string
	content string
	userID  string
}

func (v taskViewModel) ID() string {
	return v.id
}
func (v taskViewModel) Title() string {
	return v.title
}
func (v taskViewModel) Content() string {
	return v.content
}
func (v taskViewModel) UserID() string {
	return v.userID
}

type TaskController struct{}

func (c TaskController) CreateTask(title string, content string, userID string) (*taskViewModel, error) {
	input := usecase.NewCreateTaskInputData(title, content, userID)
	taskRepository := repository.NewTaskRepository()
	interactor := usecase.NewCreateTaskInteractor(input, taskRepository)
	output, err := interactor.Handle()
	viewModel := &taskViewModel{
		id:      output.ID(),
		title:   output.Title(),
		content: output.Content(),
		userID:  output.UserID(),
	}
	return viewModel, err
}

func (c TaskController) FindTask(id string) (*taskViewModel, error) {
	input := usecase.NewFindTaskInputData(id)
	taskRepository := repository.NewTaskRepository()
	interactor := usecase.NewFindTaskInteractor(input, taskRepository)
	output, err := interactor.Handle()
	viewModel := &taskViewModel{
		id:      output.ID(),
		title:   output.Title(),
		content: output.Content(),
		userID:  output.UserID(),
	}
	return viewModel, err
}
