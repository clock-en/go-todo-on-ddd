package dto

import "github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"

type CreateTaskOutputData struct {
	task taskData
}

func NewCreateTaskOutputData(task entity.Task) *CreateTaskOutputData {
	return &CreateTaskOutputData{
		task: createTaskData(task),
	}
}
func (o CreateTaskOutputData) Task() taskData {
	return o.task
}

type FetchTasksOutputData struct {
	tasks []taskData
}

func NewFetchTasksOutputData(tasks entity.Tasks) *FetchTasksOutputData {
	slice := []taskData{}
	for _, task := range tasks {
		slice = append(slice, createTaskData(task))
	}
	return &FetchTasksOutputData{tasks: slice}
}
func (o FetchTasksOutputData) Tasks() []taskData {
	return o.tasks
}

type FindTaskOutputData struct {
	task taskData
}

func NewFindTaskOutputData(task entity.Task) *FindTaskOutputData {
	return &FindTaskOutputData{
		task: createTaskData(task),
	}
}
func (o FindTaskOutputData) Task() taskData {
	return o.task
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
func createTaskData(task entity.Task) taskData {
	return taskData{
		task.ID().Value(),
		task.Title().Value(),
		task.Content().Value(),
		task.UserID().Value(),
	}
}
