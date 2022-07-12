package entity

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

// Task represents entity
type Task struct {
	id      vo.Id
	title   vo.TaskTitle
	content vo.TaskContent
	userId  vo.Id
}

type Tasks []Task

func NewTask(id vo.Id, title vo.TaskTitle, content vo.TaskContent, userId vo.Id) *Task {
	return &Task{id: id, title: title, content: content, userId: userId}
}

func (t Task) Id() vo.Id {
	return t.id
}

func (t Task) Title() vo.TaskTitle {
	return t.title
}

func (t Task) Content() vo.TaskContent {
	return t.content
}

func (t Task) UserId() vo.Id {
	return t.userId
}
