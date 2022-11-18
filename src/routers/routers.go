package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql"
	"github.com/wildanfaz/backend_skyshi/src/modules/v1/activity"
	"github.com/wildanfaz/backend_skyshi/src/modules/v1/todo"
)

func New() (*echo.Echo, error) {
	e := echo.New()

	db, err := dbmysql.New()

	activity.New(e, db)
	todo.New(e, db)

	if err != nil {
		return nil, err
	}

	return e, nil
}
