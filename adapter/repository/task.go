package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type taskRepository struct {
	dao ITaskDao
}

func NewTaskRepository(dao ITaskDao) taskRepository {
	return taskRepository{dao: dao}
}

func (r taskRepository) CreateTask(task entity.Task) error {
	return r.dao.CreateTask(
		task.ID().Value(),
		task.Title().Value(),
		task.Content().Value(),
		task.UserID().Value(),
	)
}

func (r taskRepository) FindTaskByID(id vo.ID) (*entity.Task, error) {
	row, err := r.dao.FindTaskByID(id.Value())
	if err != nil {
		return nil, err
	}
	task := createTaskEntity(row)
	return &task, nil
}

func (r taskRepository) FetchTasksByUserID(userID vo.ID) (entity.Tasks, error) {
	rows, err := r.dao.FetchTasksByUserID(userID.Value())
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
