package dao

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	_ "github.com/go-sql-driver/mysql"
)

type IUserDao interface {
	CreateUser()
	Close()
}

type UserDao struct {
	sqlHandler *SqlHandler
	IUserDao
}

func (t UserDao) CreateUser(user entity.User) (entity.User, error) {
	const sql = "INSERT INTO users (name,email,password) VALUES (?, ?, ?);"
	_, err := t.sqlHandler.connect.Exec(sql, user.Name(), user.Email(), user.Password())
	return user, err
}

func (t UserDao) Close() {
	t.sqlHandler.connect.Close()
}

func NewUserDao() *UserDao {
	sqlHandler := newSqlHandler()
	return &UserDao{sqlHandler: sqlHandler}
}
