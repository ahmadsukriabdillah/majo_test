package models

import (
	"time"
)

type User struct {
	Id        uint      `db:'id'`
	Name      string    `db:'name'`
	Username  string    `db:'user_name'`
	Password  string    `db:'password'`
	CreatedAt time.Time `db:'created_at'`
	CreatedBy uint      `db:'created_by'`
	UpdatedAt time.Time `db:'updated_at'`
	UpdatedBy uint      `db:'updated_by'`
}
