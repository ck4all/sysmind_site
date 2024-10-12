package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		//get token from request
		token := ctx.Request().Header("Authorization", "")

		if token != "" {
			_, err := facades.Auth(ctx).Parse(token)

			if err != nil {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
					"code":    http.StatusUnauthorized,
					"message": "Token Issue...!",
					"data":    nil,
					"error":   err.Error(),
				})
			} else {
				ctx.Request().Next()
			}
		} else {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"code":    http.StatusUnauthorized,
				"message": "Token Issue...!",
				"data":    token,
				"error":   "Token required..!",
			})
		}
	}
}
