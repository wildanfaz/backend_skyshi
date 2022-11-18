package interfaces

import (
	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql/models"
	"github.com/wildanfaz/backend_skyshi/src/libs"
)

type ActivityRepo interface {
	GetAllRepo() (*models.Activities, error)
	GetOneRepo(id int) (*models.Activity, error)
	CreateRepo(body *models.Activity) (*models.Activity, error)
	DeleteRepo(id int) error
	UpdateRepo(id int, body *models.Activity) (*models.Activity, error)
}

type ActivityService interface {
	GetAll() *libs.Resp
	GetOne(id int) *libs.Resp
	Create(body *models.Activity) *libs.Resp
	Delete(id int) *libs.Resp
	Update(id int, body *models.Activity) *libs.Resp
}
