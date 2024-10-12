package models

import (
	"github.com/goravel/framework/database/orm"
)

type Category struct {
	orm.Model
	Uuid     string
	Name     string
	SlugName string
	Urutan   int8
	UserUuid string
	orm.SoftDeletes
}

func (c *Category) TableName() string {
	return "categories"
}
