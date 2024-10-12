package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthChangePasswordRequest struct {
	Password string `form:"password" json:"password"`
}

func (r *AuthChangePasswordRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthChangePasswordRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"password": "required",
	}
}

func (r *AuthChangePasswordRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthChangePasswordRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthChangePasswordRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
