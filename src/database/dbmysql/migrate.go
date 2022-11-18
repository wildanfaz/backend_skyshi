package dbmysql

import (
	"context"
	"database/sql"
)

func MigUp(db *sql.DB) error {
	query := `CREATE TABLE todo4.activities (
		id BIGINT auto_increment NOT NULL,
		email varchar(100) NOT NULL,
		title varchar(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP NULL,
		CONSTRAINT activity_PK PRIMARY KEY (id)
	)
	ENGINE=InnoDB
	DEFAULT CHARSET=utf8mb4
	COLLATE=utf8mb4_0900_ai_ci;
	
	CREATE TABLE todo4.todos (
		id BIGINT auto_increment NOT NULL,
		activity_group_id varchar(100) NOT NULL,
		title varchar(100) NOT NULL,
		is_active varchar(100) NOT NULL,
		priority varchar(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP NULL,
		CONSTRAINT todo_PK PRIMARY KEY (id)
	)
	ENGINE=InnoDB
	DEFAULT CHARSET=utf8mb4
	COLLATE=utf8mb4_0900_ai_ci;`

	if _, err := db.ExecContext(context.Background(), query); err != nil {
		return err
	}

	return nil
}