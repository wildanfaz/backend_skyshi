package todo

import (
	"fmt"

	"github.com/wildanfaz/backend_skyshi/src/database/dbmysql/models"
	"github.com/wildanfaz/backend_skyshi/src/interfaces"
	"github.com/wildanfaz/backend_skyshi/src/libs"
)

type todo_service struct {
	repo interfaces.TodoRepo
}

func NewService(repo interfaces.TodoRepo) *todo_service {
	return &todo_service{repo}
}

// ** for return {} as expected output in postman documentation
var null = make(map[string]string)

func (svc *todo_service) GetAll(act int) *libs.Resp {
	data, err := svc.repo.GetAllRepo(act)

	if err != nil {
		return libs.Response(null, "Bad Request", err.Error(), 400)
	} else if len(*data) == 0 {
		return libs.Response(null, "Not Found", "Data Not Found", 404)
	}

	return libs.Response(data, "Success", "Success", 200)
}

func (svc *todo_service) GetOne(id int) *libs.Resp {
	data, err := svc.repo.GetOneRepo(id)

	if err != nil {
		return libs.Response(null, "Bad Request", err.Error(), 400)
	} else if data.Id == 0 {
		return libs.Response(null, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id), 404)
	}

	return libs.Response(data, "Success", "Success", 200)
}

func (svc *todo_service) Create(body *models.Todo) *libs.Resp {
	data, err := svc.repo.CreateRepo(body)

	var dataCreate models.CreateTodo

	if data != nil {
		dataCreate.Created_at = data.Created_at
		dataCreate.Updated_at = data.Updated_at
		dataCreate.Id = data.Id
		dataCreate.Title = data.Title
		dataCreate.Activity_group_id = data.Activity_group_id
		dataCreate.Is_Active = data.Is_Active
		dataCreate.Priority = data.Priority
	}

	if err != nil {
		return libs.Response(null, "Bad Request", err.Error(), 400)
	} else if data.Id == 0 {
		return libs.Response(null, "Not Found", "Data Not Found", 404)
	}

	return libs.Response(dataCreate, "Success", "Success", 201)
}

func (svc *todo_service) Delete(id int) *libs.Resp {
	err := svc.repo.DeleteRepo(id)

	if err != nil && err.Error() == "Not Found" {
		return libs.Response(null, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id), 404)
	} else if err != nil {
		return libs.Response(null, "Bad Request", err.Error(), 400)
	}

	return libs.Response(null, "Success", "Success", 200)
}

func (svc *todo_service) Update(id int, body *models.Todo) *libs.Resp {
	data, err := svc.repo.UpdateRepo(id, body)

	if err != nil && err.Error() == "Not Found" {
		return libs.Response(null, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id), 404)
	} else if err != nil {
		return libs.Response(null, "Bad Request", err.Error(), 400)
	}

	return libs.Response(data, "Success", "Success", 200)
}
