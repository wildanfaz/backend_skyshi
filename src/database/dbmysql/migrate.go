package dbmysql

import (
	"context"
	"database/sql"
)

func Activities(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS activities (
		id BIGINT auto_increment NOT NULL,
		email varchar(100) NOT NULL,
		title varchar(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP NULL,
		CONSTRAINT activities_PK PRIMARY KEY (id)
	)
	ENGINE=InnoDB
	`

	if _, err := db.ExecContext(context.Background(), query); err != nil {
		return err
	}

	return nil
}

func Todos(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id BIGINT auto_increment NOT NULL,
		activity_group_id integer NOT NULL,
		title varchar(100) NULL,
		is_active bool NULL,
		priority varchar(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP NULL,
		CONSTRAINT todos_PK PRIMARY KEY (id)
	)
	ENGINE=InnoDB
	`

	if _, err := db.ExecContext(context.Background(), query); err != nil {
		return err
	}

	return nil
}
