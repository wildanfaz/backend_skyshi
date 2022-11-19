package models

import "time"

type Todo struct {
	Id                int         `json:"id,omitempty"`
	Activity_group_id int         `json:"activity_group_id,omitempty"`
	Title             string      `json:"title,omitempty"`
	Is_Active         bool        `json:"is_active,omitempty"`
	Priority          string      `json:"priority,omitempty"`
	Status            int         `json:"-"`
	Created_at        time.Time   `json:"created_at"`
	Updated_at        time.Time   `json:"updated_at"`
	Deleted_at        interface{} `json:"deleted_at"`
}

type Todos []Todo

type CreateTodo struct {
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
	Id                int       `json:"id,omitempty"`
	Title             string    `json:"title,omitempty"`
	Activity_group_id int       `json:"activity_group_id,omitempty"`
	Is_Active         bool      `json:"is_active,omitempty"`
	Priority          string    `json:"priority,omitempty"`
	Status            int       `json:"-"`
}
