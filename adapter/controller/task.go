package controller

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/repository"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase/dto"
)

type taskController struct {
	dao repository.ITaskDao
}

func NewTaskController(dao repository.ITaskDao) taskController {
	return taskController{dao: dao}
}
func (c taskController) CreateTask(title string, content string, userID string) (*dto.CreateTaskOutputData, error) {
	input := dto.NewCreateTaskInputData(title, content, userID)
	taskRepository := repository.NewTaskRepository(c.dao)
	interactor := usecase.NewCreateTaskInteractor(input, taskRepository)
	return interactor.Handle()
}
func (c taskController) FindTask(id string) (*dto.FindTaskOutputData, error) {
	input := dto.NewFindTaskInputData(id)
	taskRepository := repository.NewTaskRepository(c.dao)
	interactor := usecase.NewFindTaskInteractor(input, taskRepository)
	return interactor.Handle()
}
func (c taskController) FetchUserTasks(userID string) (*dto.FetchTasksOutputData, error) {
	input := dto.NewFetchTasksInputData(userID)
	taskRepository := repository.NewTaskRepository(c.dao)
	interactor := usecase.NewFetchTasksInteractor(input, taskRepository)
	return interactor.Handle()
}

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
