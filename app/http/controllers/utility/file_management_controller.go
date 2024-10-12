package utility

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
)

type FileManagementController struct {
	//Dependent services
	FileManagementService *services.FileManagementService
}

func NewFileManagementController(fileManagementService *services.FileManagementService) *FileManagementController {
	return &FileManagementController{
		//Inject services
		FileManagementService: fileManagementService,
	}
}

func (r *FileManagementController) Index(ctx http.Context) http.Response {
	sortBy := ctx.Request().Input("sortBy[0][key]", "")
	sortMode := ctx.Request().Input("sortBy[0][order]", "")
	keyWord := ctx.Request().Input("search", "")

	itemsPerPage := ctx.Request().InputInt("itemsPerPage")
	page := ctx.Request().InputInt("page")

	result, meta, err := r.FileManagementService.Fetch(page, itemsPerPage, keyWord, sortBy, sortMode)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, err)
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"data": result,
		"meta": meta,
	})
}

func (r *FileManagementController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *FileManagementController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *FileManagementController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *FileManagementController) Destroy(ctx http.Context) http.Response {
	return nil
}
