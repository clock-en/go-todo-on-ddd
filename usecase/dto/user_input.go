package dto

type RegisterUserInputData struct {
	name     string
	email    string
	password string
}

func NewRegisterUserInputData(name string, email string, password string) RegisterUserInputData {
	return RegisterUserInputData{name: name, email: email, password: password}
}
func (i RegisterUserInputData) Name() string {
	return i.name
}
func (i RegisterUserInputData) Email() string {
	return i.email
}
func (i RegisterUserInputData) Password() string {
	return i.password
}
