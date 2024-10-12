package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthRegisterRequest struct {
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Phone string `form:"phone" json:"phone"`
}

func (r *AuthRegisterRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthRegisterRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":  "required",
		"email": "required|email",
		"phone": "required",
	}
}

func (r *AuthRegisterRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":  "Nama tidak boleh kosong",
		"email.required": "Email tidak boleh kosong",
		"phone.required": "Telpn tidak boleh kosong",
	}
}

func (r *AuthRegisterRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthRegisterRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
