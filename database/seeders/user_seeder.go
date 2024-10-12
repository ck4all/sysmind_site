package seeders

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	hassedPassword, _ := helpers.HashPassword("rahasia")

	user := models.User{
		Uuid:     uuid.New().String(),
		Name:     "Sender Technology Indonesia",
		Email:    "admin@mode.com",
		Password: hassedPassword,
		Authent:  "superadmin",
		Avatar:   "",
		Phone:    "081280008580",
		Status:   true,
	}

	return facades.Orm().Query().Save(&user)
}
