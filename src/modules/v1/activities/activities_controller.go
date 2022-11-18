package activities

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql/models"
	"github.com/wildanfaz/backend_skyshi/src/interfaces"
	"github.com/wildanfaz/backend_skyshi/src/libs"
)

type activity_ctrl struct {
	svc interfaces.ActivityService
}

func NewCtrl(svc interfaces.ActivityService) *activity_ctrl {
	return &activity_ctrl{svc}
}

func (ctrl *activity_ctrl) GetAll(c echo.Context) error {
	res := ctrl.svc.GetAll()

	return res.Send(c)
}

func (ctrl *activity_ctrl) GetOne(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		id = 0
	}

	res := ctrl.svc.GetOne(id)

	return res.Send(c)
}

func (ctrl *activity_ctrl) Create(c echo.Context) error {
	var activity models.Activity

	if err := c.Bind(&activity); err != nil {
		libs.Response(null, "Bad Request", err.Error(), 400).Send(c)
	}

	res := ctrl.svc.Create(&activity)

	return res.Send(c)
}

func (ctrl *activity_ctrl) Delete(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		id = 0
	}

	res := ctrl.svc.Delete(id)

	return res.Send(c)
}

func (ctrl *activity_ctrl) Update(c echo.Context) error {
	var activity models.Activity
	if err := c.Bind(&activity); err != nil {
		libs.Response(null, "Bad Request", err.Error(), 400).Send(c)
	}
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		id = 0
	}

	res := ctrl.svc.Update(id, &activity)

	return res.Send(c)
}
