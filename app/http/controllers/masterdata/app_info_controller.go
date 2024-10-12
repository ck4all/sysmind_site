package masterdata

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/lang"
	"goravel/app/services"
)

type AppInfoController struct {
	//Dependent services
	AppInfoService *services.AppInfoService
}

func NewAppInfoController(appInfoService *services.AppInfoService) *AppInfoController {
	return &AppInfoController{
		//Inject services
		AppInfoService: appInfoService,
	}
}

func (r *AppInfoController) Index(ctx http.Context) http.Response {
	result, err := r.AppInfoService.Show()

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, helpers.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: lang.MsgError,
			Errors:  err,
		})
	}

	return ctx.Response().Status(http.StatusOK).Json(result)
}

func (r *AppInfoController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")
	var req requests.AppInfoRequest

	validate, err := ctx.Request().ValidateRequest(&req)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, helpers.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: lang.MsgError,
			Errors:  err,
		})
	}

	if validate != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, helpers.ValidateResponse{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: lang.MsgError,
			Errors:  validate.All(),
		})
	}

	_, err = r.AppInfoService.Update(req, id)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, helpers.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: lang.MsgError,
			Errors:  err,
		})
	}

	return ctx.Response().Status(http.StatusOK).Json(http.Json{
		"code":    http.StatusOK,
		"success": true,
		"message": lang.MsgAppInfoUpdateSuccess,
	})

	return ctx.Response().Json(http.StatusOK, helpers.SuccessResponse{
		Status:  http.StatusOK,
		Success: true,
		Message: lang.MsgAppInfoUpdateSuccess,
	})
}
