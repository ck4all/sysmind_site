package services

import (
	"fmt"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type AppInfoService struct {
}

func NewAppInfoService() *AppInfoService {
	return &AppInfoService{}
}

/**
Show App Info
*/

type ShowResponse struct {
	Id          string `json:"id"`
	AppName     string `json:"app_name"`
	AppVer      string `json:"app_ver"`
	AppDesc     string `json:"app_desc"`
	AppLogo     string `json:"app_logo"`
	AppLogoPath string `json:"app_logo_path"`
	AppTheme    string `json:"app_theme"`
	AppColor    string `json:"app_color"`
	AppCompany  string `json:"app_company"`
	AppSlogan   string `json:"app_slogan"`
	AppAddress  string `json:"app_address"`
	AppWebsite  string `json:"app_website"`
	AppEmail    string `json:"app_email"`
	AppPhone    string `json:"app_phone"`
}

func (a *AppInfoService) Show() (ShowResponse, error) {
	var model models.AppInfo

	err := facades.Orm().Query().Order("id asc").First(&model)

	if err != nil {
		return ShowResponse{}, err
	}

	//filePath
	config := facades.Config()
	basePath := config.Env("APP_URL")

	filePath := fmt.Sprintf("%s/%s/%s", basePath, "api/v1/media/show-file/logo", model.AppLogo)

	result := ShowResponse{
		Id:          model.Uuid,
		AppName:     model.AppName,
		AppVer:      model.AppVer,
		AppDesc:     model.AppDesc,
		AppTheme:    model.AppTheme,
		AppColor:    model.AppColor,
		AppCompany:  model.AppCompany,
		AppSlogan:   model.AppSlogan,
		AppAddress:  model.AppAddress,
		AppWebsite:  model.AppWebsite,
		AppEmail:    model.AppEmail,
		AppPhone:    model.AppPhone,
		AppLogo:     model.AppLogo,
		AppLogoPath: filePath,
	}

	return result, nil
}

func (a *AppInfoService) Update(req requests.AppInfoRequest, id string) (models.AppInfo, error) {
	var model models.AppInfo

	err := facades.Orm().Query().Where("uuid", id).First(&model)

	if err != nil {
		return models.AppInfo{}, err
	}

	model.AppName = req.AppName
	model.AppVer = req.AppVer
	model.AppDesc = req.AppDesc
	model.AppLogo = req.AppLogo
	model.AppTheme = req.AppTheme
	model.AppColor = req.AppColor
	model.AppCompany = req.AppCompany
	model.AppSlogan = req.AppSlogan
	model.AppAddress = req.AppAddress
	model.AppWebsite = req.AppWebsite
	model.AppEmail = req.AppEmail
	model.AppPhone = req.AppPhone

	err = facades.Orm().Query().Save(&model)

	if err != nil {
		return models.AppInfo{}, err
	}

	return model, nil
}
