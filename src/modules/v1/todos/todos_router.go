package todos
import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo, db *sql.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	e.GET("/todo-items", ctrl.GetAll)
	e.GET("/todo-items/:id", ctrl.GetOne)

	e.POST("/todo-items", ctrl.Create)

	e.DELETE("/todo-items/:id", ctrl.Delete)

	e.PATCH("todo-items/:id", ctrl.Update)
}
