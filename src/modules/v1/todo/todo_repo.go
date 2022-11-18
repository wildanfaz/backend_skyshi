package todo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql/models"
)

type todo_repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *todo_repo {
	return &todo_repo{db}
}

func (repo *todo_repo) GetAllRepo(act int) (*models.Todos, error) {
	var todos models.Todos

	query := `SELECT * FROM todo WHERE activity_group_id = ?`

	rows, err := repo.db.QueryContext(context.Background(), query, act)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo models.Todo

		rows.Scan(&todo.Id, &todo.Activity_group_id, &todo.Title, &todo.Is_Active, &todo.Priority, &todo.Created_at, &todo.Updated_at, &todo.Deleted_at)

		todos = append(todos, todo)
	}

	return &todos, nil
}

func (repo *todo_repo) GetOneRepo(id int) (*models.Todo, error) {
	var todo models.Todo
	query := `SELECT * FROM todo WHERE id = ?`

	rows, err := repo.db.QueryContext(context.Background(), query, id)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&todo.Id, &todo.Activity_group_id, &todo.Title, &todo.Is_Active, &todo.Priority, &todo.Created_at, &todo.Updated_at, &todo.Deleted_at)
	}

	return &todo, nil
}

func (repo *todo_repo) CreateRepo(body *models.Todo) (*models.Todo, error) {
	var todo models.Todo

	tx, err := repo.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	query1 := `INSERT INTO todo(activity_group_id, title, is_active, priority) VALUES(?, ?, ?, ?)`

	res, err := tx.ExecContext(context.Background(), query1, body.Activity_group_id, body.Title, body.Is_Active, body.Priority)

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	query2 := `SELECT * FROM todo WHERE id = ?`

	rows1, err := tx.QueryContext(context.Background(), query2, id)

	defer rows1.Close()

	if err != nil {
		return nil, err
	}

	for rows1.Next() {
		rows1.Scan(&todo.Id, &todo.Activity_group_id, &todo.Title, &todo.Is_Active, &todo.Priority, &todo.Created_at, &todo.Updated_at, &todo.Deleted_at)
	}

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &todo, nil
}

func (repo *todo_repo) DeleteRepo(id int) error {
	query := `DELETE FROM todo WHERE id = ?`

	res, err := repo.db.ExecContext(context.Background(), query, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("Not Found")
	}

	return nil
}

func (repo *todo_repo) UpdateRepo(id int, body *models.Todo) (*models.Todo, error) {
	var todo models.Todo
	tx, err := repo.db.BeginTx(context.Background(), nil)

	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	query1 := "SELECT id FROM todo WHERE id = ?"

	rows1, err := tx.QueryContext(context.Background(), query1, id)

	if err != nil {
		return nil, err
	}

	if !rows1.Next() {
		return nil, errors.New("Not Found")
	}

	if err := rows1.Close(); err != nil {
		return nil, err
	}

	query2 := `
	UPDATE todo 
	SET 
	activity_group_id = COALESCE(nullif(?, ''), activity_group_id),
	title = COALESCE(nullif(?, ''), title),
	is_active = COALESCE(nullif(?, ''), is_active),
	priority = COALESCE(nullif(?, ''), priority)
	WHERE id = ?`

	res2, err := tx.ExecContext(context.Background(), query2, body.Activity_group_id, body.Title, body.Is_Active, body.Priority, id)

	if err != nil {
		return nil, err
	}

	rows2, err := res2.RowsAffected()

	if err != nil {
		return nil, err
	} else if rows2 == 0 {
		return nil, errors.New("No Rows Affected")
	}

	query3 := `SELECT * FROM todo WHERE id = ?`

	rows3, err := tx.QueryContext(context.Background(), query3, id)

	defer rows3.Close()

	if err != nil {
		return nil, err
	}

	for rows3.Next() {
		rows3.Scan(&todo.Id, &todo.Activity_group_id, &todo.Title, &todo.Is_Active, &todo.Priority, &todo.Created_at, &todo.Updated_at, &todo.Deleted_at)
	}

	tx.Commit()

	return &todo, nil
}
