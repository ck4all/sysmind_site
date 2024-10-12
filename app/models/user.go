package models

import (
	"database/sql"
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Uuid             string
	Name             string
	Email            string
	Password         string
	Authent          string
	Avatar           string
	Phone            string
	GoogleId         string
	RememberToken    string
	EmailVerifiedAt  sql.NullTime
	VerificationCode string
	Status           bool
	orm.SoftDeletes
}

func (u *User) TableName() string {
	return "users"
}
