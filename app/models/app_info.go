package models

import (
	"github.com/goravel/framework/database/orm"
)

type AppInfo struct {
	orm.Model
	Uuid       string
	AppName    string
	AppVer     string
	AppDesc    string
	AppLogo    string
	AppTheme   string
	AppColor   string
	AppCompany string
	AppSlogan  string
	AppAddress string
	AppWebsite string
	AppPhone   string
	AppEmail   string
	AppTw      string
	AppFb      string
	AppIg      string
	AppIn      string
	orm.SoftDeletes
}

func (a *AppInfo) TableName() string {
	return "app_infos"
}
