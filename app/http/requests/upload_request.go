package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UploadRequest struct {
	Doctype string `form:"doctype" json:"doctype"`
}

func (r *UploadRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UploadRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"doctype": "required",
	}
}

func (r *UploadRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UploadRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UploadRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
