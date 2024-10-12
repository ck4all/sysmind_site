package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
)

type MenuController struct {
	//Dependent services
	MenuService *services.MenuService
}

func NewMenuController(menuService *services.MenuService) *MenuController {
	return &MenuController{
		//Inject services
		MenuService: menuService,
	}
}

func (r *MenuController) Index(ctx http.Context) http.Response {
	result, err := r.MenuService.Fetch(ctx)

	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(result)
	}

	return ctx.Response().Status(http.StatusOK).Json(result)
}
