package utility

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/http/requests"
	"goravel/app/lang"
	"goravel/app/services"
)

type UserController struct {
	//Dependent services
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		//Inject services
		UserService: userService,
	}
}

func (r *UserController) Index(ctx http.Context) http.Response {
	sortBy := ctx.Request().Input("sortBy[0][key]", "")
	sortMode := ctx.Request().Input("sortBy[0][order]", "")
	keyWord := ctx.Request().Input("search", "")

	itemsPerPage := ctx.Request().InputInt("itemsPerPage", 10)
	page := ctx.Request().InputInt("page", 1)
	result, meta, err := r.UserService.Fetch(page, itemsPerPage, keyWord, sortBy, sortMode)

	if err != nil {
		return ctx.Response().Status(500).Json(err)
	}

	return ctx.Response().Status(200).Json(http.Json{
		"data": result,
		"meta": meta,
	})
}

func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")

	result, err := r.UserService.Show(id)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
			"errors":  err,
		})
	}

	return ctx.Response().Json(http.StatusOK, result)
}

func (r *UserController) Store(ctx http.Context) http.Response {
	var req requests.UserRequest

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

	//chek user avaliable
	checkUser, _ := r.UserService.CheckUser(req.Email)

	if checkUser {
		return ctx.Response().Json(http.StatusOK, http.Json{
			"code":    http.StatusOK,
			"success": true,
			"message": lang.MsgEmailRegistered,
		})
	}

	//register user
	result, errstore := r.UserService.Store(req)

	if errstore != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
			"errors":  errstore,
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    http.StatusOK,
		"success": true,
		"message": lang.MsgStoreSuccess,
		"data":    result,
	})
}

func (r *UserController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")
	var req requests.UserRequest

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
			"errrors": validate.All(),
		})
	}

	result, errupdate := r.UserService.Update(req, id)

	if errupdate != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
			"errors":  errupdate,
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    http.StatusOK,
		"success": true,
		"message": lang.MsgUpdateSuccess,
		"data":    result,
	})
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")

	result, err := r.UserService.Delete(id)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
			"errors":  err,
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    http.StatusOK,
		"success": true,
		"message": lang.MsgDeleteSuccess,
		"data":    result,
	})

}

func (r *UserController) UpdateProfile(ctx http.Context) http.Response {
	var req requests.UserUpdateProfilRequest

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

	//update profile proccess
	_, errupdateprofile := r.UserService.UpdateProfile(ctx, req)

	if errupdateprofile != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"code":    http.StatusOK,
		"success": true,
		"message": lang.MsgUserUpdateProfileSuccess,
	})
}
