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
func (c taskController) CreateTask(title string, content string, userID string) (*taskViewModel, error) {
	input := dto.NewCreateTaskInputData(title, content, userID)
	taskRepository := repository.NewTaskRepository(c.dao)
	interactor := usecase.NewCreateTaskInteractor(input, taskRepository)
	output, err := interactor.Handle()
	viewModel := &taskViewModel{
		id:      output.Task().ID(),
		title:   output.Task().Title(),
		content: output.Task().Content(),
		userID:  output.Task().UserID(),
	}
	return viewModel, err
}
func (c taskController) FindTask(id string) (*taskViewModel, error) {
	input := dto.NewFindTaskInputData(id)
	taskRepository := repository.NewTaskRepository(c.dao)
	interactor := usecase.NewFindTaskInteractor(input, taskRepository)
	output, err := interactor.Handle()
	viewModel := &taskViewModel{
		id:      output.Task().ID(),
		title:   output.Task().Title(),
		content: output.Task().Content(),
		userID:  output.Task().UserID(),
	}
	return viewModel, err
}
func (c taskController) FetchUserTasks(userID string) ([]taskViewModel, error) {
	input := dto.NewFetchTasksInputData(userID)
	taskRepository := repository.NewTaskRepository(c.dao)
	interactor := usecase.NewFetchTasksInteractor(input, taskRepository)
	output, err := interactor.Handle()
	viewModel := []taskViewModel{}
	for _, task := range output.Tasks() {
		viewModel = append(viewModel, taskViewModel{
			id:      task.ID(),
			title:   task.Title(),
			content: task.Content(),
			userID:  task.UserID(),
		})
	}
	return viewModel, err
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
