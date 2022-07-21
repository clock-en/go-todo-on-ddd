package entity

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type Task struct {
	id      vo.ID
	title   vo.TaskTitle
	content vo.TaskContent
	userID  vo.ID
}

type Tasks []Task

func NewTask(id vo.ID, title vo.TaskTitle, content vo.TaskContent, userID vo.ID) *Task {
	return &Task{id: id, title: title, content: content, userID: userID}
}

func (t Task) ID() vo.ID {
	return t.id
}

func (t Task) Title() vo.TaskTitle {
	return t.title
}

func (t Task) Content() vo.TaskContent {
	return t.content
}

func (t Task) UserID() vo.ID {
	return t.userID
}
