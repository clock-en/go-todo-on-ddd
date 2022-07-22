package controller

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase"
)

type taskController struct{}

func (c taskController) create(title string, content string, userID string) (model.Task, error) {
	inputData := usecase.NewCreateTaskInput(title, content, userID)
	interactor := usecase.NewCreateTaskUsecase(inputData)
	output, err := interactor.Handler()
	taskDataModel := model.Task{
		ID:      output.Task().ID().Value(),
		Title:   output.Task().Title().Value(),
		Content: output.Task().Content().Value(),
		UserID:  output.Task().UserID().Value(),
	}
	return taskDataModel, err
}
