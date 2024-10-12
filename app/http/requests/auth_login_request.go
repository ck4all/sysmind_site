package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthLoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *AuthLoginRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthLoginRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email",
		"password": "required",
	}
}

func (r *AuthLoginRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"email.required":    "Email tidak boleh kosong",
		"password.required": "Kata sandi tidak boleh kosong",
	}
}

func (r *AuthLoginRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthLoginRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
