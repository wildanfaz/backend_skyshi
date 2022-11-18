package interfaces

import (
	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql/models"
	"github.com/wildanfaz/backend_skyshi/src/libs"
)

type TodoRepo interface {
	GetAllRepo(act int) (*models.Todos, error)
	GetOneRepo(id int) (*models.Todo, error)
	CreateRepo(body *models.Todo) (*models.Todo, error)
	DeleteRepo(id int) error
	UpdateRepo(id int, body *models.Todo) (*models.Todo, error)
}

type TodoService interface {
	GetAll(act int) *libs.Resp
	GetOne(id int) *libs.Resp
	Create(body *models.Todo) *libs.Resp
	Delete(id int) *libs.Resp
	Update(id int, body *models.Todo) *libs.Resp
}
