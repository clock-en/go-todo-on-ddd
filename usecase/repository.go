package usecase

type ITaskRepository interface {
	Create(id string, title string, content string, userID string) error
}
