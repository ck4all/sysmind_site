package seeders

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type AppInfoSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *AppInfoSeeder) Signature() string {
	return "AppInfoSeeder"
}

// Run executes the seeder logic.
func (s *AppInfoSeeder) Run() error {
	appinfo := models.AppInfo{
		Uuid:       uuid.New().String(),
		AppName:    "Sender Technology Template Apps",
		AppVer:     "1.00",
		AppDesc:    "Aplikasi Golang Master Template For All Apps",
		AppLogo:    "logo.png",
		AppTheme:   "light",
		AppColor:   "indigo",
		AppCompany: "Sender Technology Indonesia, PT",
		AppSlogan:  "We Are The Best",
		AppAddress: "Perum Sudirman Indah Blok H1/15",
		AppWebsite: "https://sendertechid.com",
		AppEmail:   "selamet.antsoftmedia@gmail.com",
		AppPhone:   "081280008580",
	}
	return facades.Orm().Query().Create(&appinfo)
}
