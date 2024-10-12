package helpers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

/*
*
Get User Information
*/
func GetUserInfo(ctx http.Context) (models.User, error) {
	var user models.User

	err := facades.Auth(ctx).User(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
