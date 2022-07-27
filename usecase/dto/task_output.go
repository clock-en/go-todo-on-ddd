package dto

import "github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"

type CreateTaskOutputData struct {
	task        *taskData
	inputErrors InputErrors
}

func NewCreateTaskOutputData(task *entity.Task, inputErrors InputErrors) *CreateTaskOutputData {
	if task == nil {
		return &CreateTaskOutputData{
			task:        nil,
			inputErrors: inputErrors,
		}
	}
	return &CreateTaskOutputData{
		task:        createTaskData(task),
		inputErrors: nil,
	}
}
func (o CreateTaskOutputData) Task() taskData {
	return *o.task
}
func (o CreateTaskOutputData) InputErrors() InputErrors {
	return o.inputErrors
}

type FetchTasksOutputData struct {
	tasks       []taskData
	inputErrors InputErrors
}

func NewFetchTasksOutputData(tasks entity.Tasks, inputErrors InputErrors) *FetchTasksOutputData {
	slice := []taskData{}
	for _, task := range tasks {
		slice = append(slice, *createTaskData(&task))
	}
	return &FetchTasksOutputData{tasks: slice, inputErrors: inputErrors}
}
func (o FetchTasksOutputData) Tasks() []taskData {
	return o.tasks
}
func (o FetchTasksOutputData) InputErrors() InputErrors {
	return o.inputErrors
}

type FindTaskOutputData struct {
	task        *taskData
	inputErrors InputErrors
}

func NewFindTaskOutputData(task *entity.Task, inputErrors InputErrors) *FindTaskOutputData {
	if task == nil {
		return &FindTaskOutputData{
			task:        nil,
			inputErrors: inputErrors,
		}
	}
	return &FindTaskOutputData{
		task:        createTaskData(task),
		inputErrors: nil,
	}
}
func (o FindTaskOutputData) Task() taskData {
	return *o.task
}
func (o FindTaskOutputData) InputErrors() InputErrors {
	return o.inputErrors
}

type taskData struct {
	id      string
	title   string
	content string
	userID  string
}

func (t taskData) ID() string {
	return t.id
}
func (t taskData) Title() string {
	return t.title
}
func (t taskData) Content() string {
	return t.content
}
func (t taskData) UserID() string {
	return t.userID
}
func createTaskData(task *entity.Task) *taskData {
	return &taskData{
		task.ID().Value(),
		task.Title().Value(),
		task.Content().Value(),
		task.UserID().Value(),
	}
}
