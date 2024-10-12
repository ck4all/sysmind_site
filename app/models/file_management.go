package models

import (
	"github.com/goravel/framework/database/orm"
)

type FileManagement struct {
	orm.Model
	Uuid       string
	FolderName string
	FileName   string
	Ext        string
	Size       int64
	Type       string
	Used       int
	UserUuid   string
	orm.SoftDeletes
}

func (f *FileManagement) TableName() string {
	return "file_managements"
}
