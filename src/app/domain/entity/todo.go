package entity

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

// this Todo represents entity
type Todo struct {
	id     vo.Id
	title  vo.TodoTitle
	task   vo.TodoTask
	userId vo.Id
}

type Todos []Todo

func NewTodo(id vo.Id, title vo.TodoTitle, task vo.TodoTask, userId vo.Id) *Todo {
	return &Todo{id: id, title: title, task: task, userId: userId}
}

func (t Todo) Id() vo.Id {
	return t.id
}

func (t Todo) Title() vo.TodoTitle {
	return t.title
}

func (t Todo) Task() vo.TodoTask {
	return t.task
}

func (t Todo) UserId() vo.Id {
	return t.userId
}
