package activity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql/models"
)

type activity_repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *activity_repo {
	return &activity_repo{db}
}

func (repo *activity_repo) GetAllRepo() (*models.Activities, error) {
	var activities models.Activities

	query := `SELECT * FROM activity`

	rows, err := repo.db.QueryContext(context.Background(), query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var activity models.Activity

		rows.Scan(&activity.Id, &activity.Email, &activity.Title, &activity.Created_at, &activity.Updated_at, &activity.Deleted_at)

		activities = append(activities, activity)
	}

	return &activities, nil
}

func (repo *activity_repo) GetOneRepo(id int) (*models.Activity, error) {
	var activity models.Activity
	query := `SELECT * FROM activity WHERE id = ?`

	rows, err := repo.db.QueryContext(context.Background(), query, id)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&activity.Id, &activity.Email, &activity.Title, &activity.Created_at, &activity.Updated_at, &activity.Deleted_at)
	}

	return &activity, nil
}

func (repo *activity_repo) CreateRepo(body *models.Activity) (*models.Activity, error) {
	var activity models.Activity
	tx, err := repo.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	query := `SELECT id FROM activity WHERE email = ?`

	rows1, err := tx.QueryContext(context.Background(), query, body.Email)

	defer rows1.Close()

	if err != nil {
		return nil, err
	}

	for rows1.Next() {
		rows1.Scan(&activity.Id)
		if activity.Id != 0 {
			return nil, errors.New("Email Already Exists")
		}
	}

	query2 := `INSERT INTO activity(email, title) VALUES(?, ?)`

	res, err := tx.ExecContext(context.Background(), query2, body.Email, body.Title)

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	query3 := `SELECT * FROM activity WHERE id = ?`

	rows2, err := tx.QueryContext(context.Background(), query3, id)

	defer rows2.Close()

	if err != nil {
		return nil, err
	}

	for rows2.Next() {
		rows2.Scan(&activity.Id, &activity.Email, &activity.Title, &activity.Created_at, &activity.Updated_at, &activity.Deleted_at)
	}

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &activity, nil
}

func (repo *activity_repo) DeleteRepo(id int) error {
	query := `DELETE FROM activity WHERE id = ?`

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

func (repo *activity_repo) UpdateRepo(id int, body *models.Activity) (*models.Activity, error) {
	var activity models.Activity
	tx, err := repo.db.BeginTx(context.Background(), nil)

	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	query1 := "SELECT id FROM activity WHERE id = ?"

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

	query2 := `SELECT id, email FROM activity WHERE email = ?`

	rows2, err := tx.QueryContext(context.Background(), query2, body.Email)

	defer rows2.Close()

	if err != nil {
		return nil, err
	}

	for rows2.Next() {
		rows2.Scan(&activity.Id, &activity.Email)
		if activity.Id != 0 && activity.Email != "" {
			return nil, errors.New("Email Already Exists")
		}
	}

	query3 := `
	UPDATE activity 
	SET 
	email = COALESCE(nullif(?, ''), email),
	title = COALESCE(nullif(?, ''), title)
	WHERE id = ?`

	res2, err := tx.ExecContext(context.Background(), query3, body.Email, body.Title, id)

	if err != nil {
		return nil, err
	}

	rows3, err := res2.RowsAffected()

	if err != nil {
		return nil, err
	} else if rows3 == 0 {
		return nil, errors.New("No Rows Affected")
	}

	query4 := `SELECT * FROM activity WHERE id = ?`

	rows4, err := tx.QueryContext(context.Background(), query4, id)

	defer rows4.Close()

	if err != nil {
		return nil, err
	}

	for rows4.Next() {
		rows4.Scan(&activity.Id, &activity.Email, &activity.Title, &activity.Created_at, &activity.Updated_at, &activity.Deleted_at)
	}

	tx.Commit()

	return &activity, nil
}
