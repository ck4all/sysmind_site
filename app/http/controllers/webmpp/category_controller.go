package webmpp

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
)

type CategoryController struct {
	//Dependent services
	CategoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		//Inject services
		CategoryService: categoryService,
	}
}

func (r *CategoryController) Index(ctx http.Context) http.Response {
	sortBy := ctx.Request().Input("sortBy[0][key]", "")
	sortMode := ctx.Request().Input("sortBy[0][order]", "")
	keyWord := ctx.Request().Input("search", "")

	itemsPerPage := ctx.Request().InputInt("itemsPerPage", 10) 
	page := ctx.Request().InputInt("page",1)
	result, meta, err := r.CategoryService.Fetch(page, itemsPerPage, keyWord, sortBy, sortMode)

	if err != nil {
		return ctx.Response().Status(500).Json(err)
	}

	return ctx.Response().Status(200).Json(http.Json{
		"data": result,
		"meta": meta,
	})
}

func (r *CategoryController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *CategoryController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *CategoryController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	return nil
}
