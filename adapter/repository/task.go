package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
)

// TODO: Daoへの依存を検討する (interface `ITaskDao` を作成して抽象化？)
type TaskRepository struct{}

func (r TaskRepository) Create(id string, title string, content string, userID string) error {
	dao := dao.TaskDao{}
	err := dao.Create(id, title, content, userID)
	return err
}
