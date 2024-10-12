package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AppInfoRequest struct {
	Id         string `form:"id" json:"id"`
	AppName    string `form:"app_name" json:"app_name"`
	AppVer     string `form:"app_ver" json:"app_ver"`
	AppDesc    string `form:"app_desc" json:"app_desc"`
	AppLogo    string `form:"app_logo" json:"app_logo"`
	AppTheme   string `form:"app_theme" json:"app_theme"`
	AppColor   string `form:"app_color" json:"app_color"`
	AppCompany string `form:"app_company" json:"app_company"`
	AppSlogan  string `form:"app_slogan" json:"app_slogan"`
	AppAddress string `form:"app_address" json:"app_address"`
	AppWebsite string `form:"app_website" json:"app_website"`
	AppPhone   string `form:"app_phone" json:"app_phone"`
	AppEmail   string `form:"app_email" json:"app_email"`
}

func (r *AppInfoRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AppInfoRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"id":       "required",
		"app_name": "required",
	}
}

func (r *AppInfoRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AppInfoRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AppInfoRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
