package models

import "time"

type Activity struct {
	Id         int         `json:"id,omitempty"`
	Email      string      `json:"email,omitempty"`
	Title      string      `json:"title,omitempty"`
	Created_at time.Time   `json:"created_at"`
	Updated_at time.Time   `json:"updated_at"`
	Deleted_at interface{} `json:"deleted_at"`
}

type Activities []Activity

type CreateActivity struct {
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Id         int       `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Email      string    `json:"email,omitempty"`
}
