package dao

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	host     string
	database string
	username string
	password string
	port     string
)

type SqlHandler struct {
	connect *sql.DB
}

func newSqlHandler() *SqlHandler {
	db := getConnect()
	return &SqlHandler{connect: db}
}

func getConnect() *sql.DB {
	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	host = os.Getenv("MYSQL_HOST")
	port = "3306"
	database = os.Getenv("MYSQL_DB")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	db, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	return db
}
