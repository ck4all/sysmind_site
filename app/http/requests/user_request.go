package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserRequest struct {
	Name    string `form:"name" json:"name"`
	Email   string `form:"email" json:"email"`
	Authent string `form:"authent" json:"authent"`
	Phone   string `form:"phone" json:"phone"`
	Avatar  string `form:"avatar" json:"avatar"`
	Status  bool   `form:"status" json:"status"`
}

func (r *UserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":    "required",
		"email":   "required|email",
		"authent": "required",
		"phone":   "required",
	}
}

func (r *UserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":    "Nama pengguna tidak boleh kosong..!",
		"email.required":   "Email tidak boleh kosong..!",
		"email.email":      "Penulisan email tidak benar...!",
		"authent.required": "Level pengguna wajib dipilih..!",
	}
}

func (r *UserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
