package utility

import (
	"github.com/goravel/framework/contracts/http"
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
	return nil
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	return nil
}
