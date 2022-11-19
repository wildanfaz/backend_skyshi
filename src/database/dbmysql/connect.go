package dbmysql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")
	port := os.Getenv("MYSQL_PORT")

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	db, err := sql.Open("mysql", config)

	if err != nil {
		return nil, errors.New("error connect mysql")
	}

	db.SetMaxIdleConns(500)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxIdleTime(time.Minute * 2)
	db.SetConnMaxLifetime(time.Minute * 4)

	if err := Activities(db); err != nil {
		return nil, err
	}

	if err := Todos(db); err != nil {
		return nil, err
	}

	return db, nil
}
