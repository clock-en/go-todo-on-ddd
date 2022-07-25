package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
)

type taskRepository struct{}

func NewTaskRepository() taskRepository {
	return taskRepository{}
}

func (r taskRepository) Create(task entity.Task) (*entity.Task, error) {
	dao := dao.TaskDao{}
	err := dao.Create(
		task.ID().Value(),
		task.Title().Value(),
		task.Content().Value(),
		task.UserID().Value(),
	)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (q taskRepository) FindById(id vo.ID) (*entity.Task, error) {
	dao := dao.TaskDao{}
	row, err := dao.FindById(id.Value())
	if err != nil {
		return nil, err
	}
	task := taskFactory(row)
	return &task, nil
}

type ITaskRow interface {
	ID() string
	Title() string
	Content() string
	UserID() string
}

type ITaskDao interface {
	Create(id string, title string, content string, userID string) error
	FindById(id string) (ITaskRow, error)
}

func taskFactory(taskRow ITaskRow) entity.Task {
	id, _ := vo.NewID(taskRow.ID())
	title, _ := vo.NewTaskTitle(taskRow.ID())
	content, _ := vo.NewTaskContent(taskRow.ID())
	userID, _ := vo.NewID(taskRow.ID())
	// TODO: voのエラー処理を実装
	return *entity.NewTask(
		*id,
		*title,
		*content,
		*userID,
	)
}
