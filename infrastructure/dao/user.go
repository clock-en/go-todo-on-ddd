package dao

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/repository"
	_ "github.com/go-sql-driver/mysql"
)

type UserDao struct{}

func (t UserDao) CreateUser(id string, name string, email string, password string) error {
	handler := newSqlHandler()
	defer handler.connect.Close()
	const stmt = "INSERT INTO users (id, name,email,password) VALUES (?, ?, ?, ?);"
	_, err := handler.connect.Exec(stmt, id, name, email, password)
	return err
}
func (UserDao) FindRegisteredUser(name string, email string) (repository.IUserRow, error) {
	handler := newSqlHandler()
	defer handler.connect.Close()

	user := &userRow{}
	const stmt = "SELECT * FROM users WHERE name=? OR email=?;"
	row := handler.connect.QueryRow(stmt, name, email)
	err := row.Scan(&user.id, &user.name, &user.email, &user.password, &user.createdAt, &user.updatedAt)

	return user, err
}

type userRow struct {
	id        string
	name      string
	email     string
	password  string
	createdAt string
	updatedAt string
}

func (r userRow) ID() string {
	return r.id
}
func (r userRow) Name() string {
	return r.name
}
func (r userRow) Email() string {
	return r.email
}
func (r userRow) Password() string {
	return r.password
}
