package helpers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func CheckToken(ctx http.Context) bool {
	token := ctx.Request().Header("Authorization")

	_, err := facades.Auth(ctx).Parse(token)

	if err != nil {
		return false
	}

	return true
}
