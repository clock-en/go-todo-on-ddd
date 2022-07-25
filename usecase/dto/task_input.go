package dto

type CreateTaskInputData struct {
	title   string
	content string
	userID  string
}

func NewCreateTaskInputData(title string, content string, userID string) CreateTaskInputData {
	return CreateTaskInputData{title: title, content: content, userID: userID}
}
func (i CreateTaskInputData) Title() string {
	return i.title
}
func (i CreateTaskInputData) Content() string {
	return i.content
}
func (i CreateTaskInputData) UserID() string {
	return i.userID
}

type FetchTasksInputData struct {
	userID string
}

func NewFetchTasksInputData(userID string) FetchTasksInputData {
	return FetchTasksInputData{userID: userID}
}
func (i FetchTasksInputData) UserID() string {
	return i.userID
}

type FindTaskInputData struct {
	id string
}

func NewFindTaskInputData(id string) FindTaskInputData {
	return FindTaskInputData{id: id}
}
func (i FindTaskInputData) ID() string {
	return i.id
}
