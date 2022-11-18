package todo

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql/models"
	"github.com/wildanfaz/backend_skyshi/src/interfaces"
	"github.com/wildanfaz/backend_skyshi/src/libs"
)

type todo_ctrl struct {
	svc interfaces.TodoService
}

func NewCtrl(svc interfaces.TodoService) *todo_ctrl {
	return &todo_ctrl{svc}
}

func (ctrl *todo_ctrl) GetAll(c echo.Context) error {
	qp := c.QueryParam("activity_group_id")

	id, err := strconv.Atoi(qp)

	if err != nil {
		id = 0
	}

	res := ctrl.svc.GetAll(id)

	return res.Send(c)
}

func (ctrl *todo_ctrl) GetOne(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		id = 0
	}

	res := ctrl.svc.GetOne(id)

	return res.Send(c)
}

func (ctrl *todo_ctrl) Create(c echo.Context) error {
	var todo models.Todo

	if err := c.Bind(&todo); err != nil {
		libs.Response(null, "Bad Request", err.Error(), 400).Send(c)
	}

	res := ctrl.svc.Create(&todo)

	return res.Send(c)
}

func (ctrl *todo_ctrl) Delete(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		id = 0
	}

	res := ctrl.svc.Delete(id)

	return res.Send(c)
}

func (ctrl *todo_ctrl) Update(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		libs.Response(null, "Bad Request", err.Error(), 400).Send(c)
	}
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		id = 0
	}

	res := ctrl.svc.Update(id, &todo)

	return res.Send(c)
}
