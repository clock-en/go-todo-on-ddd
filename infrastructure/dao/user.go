package dao

import (
	_ "github.com/go-sql-driver/mysql"
)

type UserDao struct {
	handler *sqlHandler
}

func (t UserDao) CreateUser(name string, email string, password string) error {
	const sql = "INSERT INTO users (name,email,password) VALUES (?, ?, ?);"
	_, err := t.handler.connect.Exec(sql, name, email, password)
	return err
}

func (t UserDao) Close() {
	t.handler.connect.Close()
}

func NewUserDao() *UserDao {
	sqlHandler := newSqlHandler()
	return &UserDao{handler: sqlHandler}
}
