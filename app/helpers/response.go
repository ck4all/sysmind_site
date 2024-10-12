package helpers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"goravel/app/lang"
)

type MetaResponse struct {
	Total       int64   `json:"total"`
	PerPage     int     `json:"per_page"`
	CurrentPage int     `json:"current_page"`
	FirstPage   int     `json:"first_page"`
	LastPage    float64 `json:"last_page"`
}

type ComboResponse struct {
	Value any    `json:"value"`
	Title string `json:"title"`
}

type StatusResponse struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  error  `json:"errors"`
}

type ValidateResponse struct {
	Status  int                          `json:"status"`
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Errors  map[string]map[string]string `json:"errors"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type FetchResponse struct {
	Data any `json:"data"`
	Meta any `json:"meta"`
}

type DeleteResponse struct {
	Uuid string `json:"id"`
}

func GetGeneralResponse(ctx http.Context, status int, success bool, message string, err error) http.Response {
	return ctx.Response().Json(status, http.Json{"status": status, "success": success, message: message, "errors": err})
}

func GetErrorResponse(ctx http.Context, err error) http.Response {
	result := ErrorResponse{
		Status:  http.StatusInternalServerError,
		Success: false,
		Message: lang.MsgError,
		Errors:  err,
	}

	return ctx.Response().Json(http.StatusInternalServerError, result)
}

func GetSuccessResponse(ctx http.Context, message string, data any) http.Response {
	result := SuccessResponse{
		Status:  http.StatusOK,
		Success: true,
		Message: message,
		Data:    data,
	}

	return ctx.Response().Json(http.StatusOK, result)
}

func GetFetchResponse(ctx http.Context, datas any, meta MetaResponse) http.Response {
	result := FetchResponse{
		Data: datas,
		Meta: meta,
	}

	return ctx.Response().Json(http.StatusOK, result)
}

func GetValidationResponse(ctx http.Context, error validation.Errors) http.Response {
	result := ValidateResponse{
		Status:  http.StatusUnprocessableEntity,
		Message: lang.MsgError,
		Errors:  error.All(),
	}

	return ctx.Response().Json(http.StatusUnprocessableEntity, result)
}
