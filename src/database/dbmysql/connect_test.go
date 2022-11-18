package dbmysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnect(t *testing.T) {
	config := fmt.Sprint("xxxx:xxxxx@tcp(127.17.0.1:3306)/todo4")

	db, err := sql.Open("mysql", config)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
