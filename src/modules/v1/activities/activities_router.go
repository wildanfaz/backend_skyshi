package activities

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo, db *sql.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	e.GET("/activity-groups", ctrl.GetAll)
	e.GET("/activity-groups/:id", ctrl.GetOne)

	e.POST("/activity-groups", ctrl.Create)

	e.DELETE("/activity-groups/:id", ctrl.Delete)

	e.PATCH("/activity-groups/:id", ctrl.Update)
}
