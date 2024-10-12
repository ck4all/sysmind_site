package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserUpdateProfilRequest struct {
	Name    string `form:"name" json:"name"`
	Email   string `form:"email" json:"email"`
	Authent string `form:"authent" json:"authent"`
	Phone   string `form:"phone" json:"phone"`
	Avatar  string `form:"avatar" json:"avatar"`
	Status  bool   `form:"status" json:"status"`
}

func (r *UserUpdateProfilRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserUpdateProfilRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name": "required",
	}
}

func (r *UserUpdateProfilRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserUpdateProfilRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserUpdateProfilRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
