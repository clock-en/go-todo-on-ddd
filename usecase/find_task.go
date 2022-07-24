package usecase

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type findTaskInputData struct {
	id string
}

func NewFindTaskInputData(id string) findTaskInputData {
	return findTaskInputData{id: id}
}

type findTaskOutputData struct {
	id      string
	title   string
	content string
	userID  string
}

func (o findTaskOutputData) ID() string {
	return o.id
}
func (o findTaskOutputData) Title() string {
	return o.title
}
func (o findTaskOutputData) Content() string {
	return o.content
}
func (o findTaskOutputData) UserID() string {
	return o.userID
}

type findTaskInteractor struct {
	input          findTaskInputData
	taskRepository ITaskRepository
}

func NewFindTaskInteractor(input createTaskInputData, taskRepository ITaskRepository) *createTaskInteractor {
	return &createTaskInteractor{input: input, taskRepository: taskRepository}
}

func (i findTaskInteractor) Handle() (*findTaskOutputData, error) {
	id, _ := vo.NewID(i.input.id)

	task, dberror := i.taskRepository.FindById(id.Value())
	if dberror != nil {
		return nil, dberror
	}

	output := &findTaskOutputData{
		id:      task.ID(),
		title:   task.Title(),
		content: task.Content(),
		userID:  task.UserID(),
	}
	return output, nil
}
