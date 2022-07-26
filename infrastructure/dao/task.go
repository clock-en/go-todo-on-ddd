package dao

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/repository"
	_ "github.com/go-sql-driver/mysql"
)

type TaskDao struct{}

func (TaskDao) CreateTask(id string, title string, content string, userID string) error {
	handler := newSqlHandler()
	defer handler.connect.Close()
	const stmt = "INSERT INTO tasks (id, title, content, user_id) VALUES (?, ?, ?, ?);"
	_, err := handler.connect.Exec(stmt, id, title, content, userID)
	return err
}
func (TaskDao) FindTaskByID(id string) (repository.ITaskRow, error) {
	handler := newSqlHandler()
	defer handler.connect.Close()

	task := &taskRow{}
	const stmt = "SELECT * FROM tasks WHERE id=?;"
	row := handler.connect.QueryRow(stmt, id)
	err := row.Scan(&task.id, &task.title, &task.content, &task.userID, &task.createdAt, &task.updatedAt)

	return task, err
}
func (TaskDao) FetchTasksByUserID(userID string) ([]repository.ITaskRow, error) {
	handler := newSqlHandler()
	defer handler.connect.Close()

	tasks := []repository.ITaskRow{}
	const stmt = "SELECT * FROM tasks WHERE user_id=?;"
	rows, err := handler.connect.Query(stmt, userID)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		task := taskRow{}
		if err := rows.Scan(&task.id, &task.title, &task.content, &task.userID, &task.createdAt, &task.updatedAt); err != nil {
			panic(err)
		}
		tasks = append(tasks, task)
	}

	return tasks, err
}

type taskRow struct {
	id        string
	title     string
	content   string
	userID    string
	createdAt string
	updatedAt string
}

func (r taskRow) ID() string {
	return r.id
}
func (r taskRow) Title() string {
	return r.title
}
func (r taskRow) Content() string {
	return r.content
}
func (r taskRow) UserID() string {
	return r.userID
}
