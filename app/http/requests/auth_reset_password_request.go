package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthResetPasswordRequest struct {
	Email string `form:"email" json:"email"`
}

func (r *AuthResetPasswordRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthResetPasswordRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email": "required|email",
	}
}

func (r *AuthResetPasswordRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"email.required": "Email tidak boleh kosong",
	}
}

func (r *AuthResetPasswordRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthResetPasswordRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
