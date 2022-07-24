package usecase

type ITaskRow interface {
	ID() string
	Title() string
	Content() string
	UserID() string
}

type ITaskRepository interface {
	Create(id string, title string, content string, userID string) error
	FindById(id string) (ITaskRow, error)
}
