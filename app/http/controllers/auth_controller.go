package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/lang"
	"goravel/app/services"
)

type AuthController struct {
	//Dependent services
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		//Inject services
		AuthService: authService,
	}
}

func (r *AuthController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *AuthController) Register(ctx http.Context) http.Response {
	var req requests.AuthRegisterRequest

	//run validation request data
	resultVal, errorVal := ctx.Request().ValidateRequest(&req)

	//catch error validation and send response error
	if errorVal != nil {
		return helpers.GetErrorResponse(ctx, errorVal)
	}

	//get response validation if request not valid
	if resultVal != nil {
		return helpers.GetValidationResponse(ctx, resultVal)
	}

	//check email user is avaliabel on database
	checkUser := r.AuthService.CheckEmail(req.Email)

	if !checkUser {
		return helpers.GetGeneralResponse(ctx, http.StatusUnprocessableEntity, false, lang.MsgEmailRegistered, nil)
	}

	result, err := r.AuthService.Register(req)

	if err != nil {
		return helpers.GetErrorResponse(ctx, err)
	}

	return helpers.GetSuccessResponse(ctx, lang.MsgRegisterUserSuccess, result)
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	var req requests.AuthLoginRequest

	//run validation data
	resultVal, errVal := ctx.Request().ValidateRequest(&req)

	if errVal != nil {
		return helpers.GetErrorResponse(ctx, errVal)
	}

	if resultVal != nil {
		return helpers.GetValidationResponse(ctx, resultVal)
	}

	result, err := r.AuthService.Login(req)

	if !err {
		return helpers.GetGeneralResponse(ctx, 401, false, lang.MsgLoginError, nil)
	}

	//get token
	token, _ := facades.Auth(ctx).Login(result)

	payload, _ := facades.Auth(ctx).Parse(token)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"status":  http.StatusOK,
		"success": true,
		"response": http.Json{
			"token": http.Json{
				"type":      "Bearer",
				"token":     token,
				"expaireAt": payload,
			},
		},
	})
}

func (r *AuthController) Reset(ctx http.Context) http.Response {
	var req requests.AuthResetPasswordRequest

	resultVal, errVal := ctx.Request().ValidateRequest(&req)

	if errVal != nil {
		return helpers.GetErrorResponse(ctx, errVal)
	}

	if resultVal != nil {
		return helpers.GetValidationResponse(ctx, resultVal)
	}

	result, err := r.AuthService.ResetPassword(req)

	if err != nil {
		return helpers.GetErrorResponse(ctx, err)
	}

	return helpers.GetSuccessResponse(ctx, lang.MsgResetPasswordSuccess, result)
}

func (r *AuthController) ChangePassword(ctx http.Context) http.Response {
	var req requests.AuthChangePasswordRequest

	validate, err := ctx.Request().ValidateRequest(&req)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
			"errors":  err,
		})
	}

	if validate != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"code":    http.StatusUnprocessableEntity,
			"success": false,
			"message": lang.MsgError,
			"errors":  validate.All(),
		})
	}

	//proses update password
	_, err = r.AuthService.ChangPassword(ctx, req)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
			"errors":  err,
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    http.StatusUnprocessableEntity,
		"success": true,
		"message": lang.MsgChangePasswordSuccess,
	})
}

func (r *AuthController) UserInfo(ctx http.Context) http.Response {
	result, err := r.AuthService.UserInfo(ctx)

	if err != nil {
		return helpers.GetErrorResponse(ctx, err)
	}

	return helpers.GetSuccessResponse(ctx, lang.MsgSuccess, result)
}

func (r *AuthController) Logout(ctx http.Context) http.Response {
	result, err := r.AuthService.Logout(ctx)

	if err != nil {
		return helpers.GetErrorResponse(ctx, err)
	}

	return helpers.GetSuccessResponse(ctx, lang.MsgSuccess, result)
}
