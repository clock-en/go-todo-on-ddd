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

func (r taskRepository) CreateTask(task entity.Task) (*entity.Task, error) {
	dao := dao.TaskDao{}
	err := dao.CreateTask(
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

func (q taskRepository) FindTaskByID(id vo.ID) (*entity.Task, error) {
	dao := dao.TaskDao{}
	row, err := dao.FindTaskByID(id.Value())
	if err != nil {
		return nil, err
	}
	task := createTaskEntity(row)
	return &task, nil
}

func (q taskRepository) FetchTasksByUserID(userID vo.ID) (entity.Tasks, error) {
	dao := dao.TaskDao{}
	rows, err := dao.FetchTasksByUserID(userID.Value())
	if err != nil {
		return nil, err
	}
	tasks := []entity.Task{}
	for _, row := range rows {
		tasks = append(tasks, createTaskEntity(row))
	}
	return tasks, nil
}

type ITaskRow interface {
	ID() string
	Title() string
	Content() string
	UserID() string
}

type ITaskDao interface {
	CreateTask(id string, title string, content string, userID string) error
	FindTaskByID(id string) (ITaskRow, error)
	FetchTasksByUserID(userID string) ([]ITaskRow, error)
}

func createTaskEntity(taskRow ITaskRow) entity.Task {
	id, _ := vo.NewID(taskRow.ID())
	title, _ := vo.NewTaskTitle(taskRow.Title())
	content, _ := vo.NewTaskContent(taskRow.Content())
	userID, _ := vo.NewID(taskRow.UserID())
	// TODO: voのエラー処理を実装
	return *entity.NewTask(
		*id,
		*title,
		*content,
		*userID,
	)
}
