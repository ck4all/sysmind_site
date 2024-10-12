package utility

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests"
	"goravel/app/lang"
	"goravel/app/services"
)

type MediaController struct {
	//Dependent services
	FileManagementService *services.FileManagementService
}

func NewMediaController(fileManagementService *services.FileManagementService) *MediaController {
	return &MediaController{
		//Inject services
		FileManagementService: fileManagementService,
	}
}

func (c *MediaController) Upload(ctx http.Context) http.Response {
	var req requests.UploadRequest

	validateerr, err := ctx.Request().ValidateRequest(&req)

	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
			"errors":  err,
		})
	}

	if validateerr != nil {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{
			"status":  http.StatusUnprocessableEntity,
			"success": false,
			"message": lang.MsgError,
			"errors":  validateerr.All(),
		})
	}

	//get file from request
	file, err := ctx.Request().File("file")

	fileName := fmt.Sprintf("%s.%s", uuid.New().String(), file.GetClientOriginalExtension())

	//get file type
	fileType, _ := file.MimeType()

	//get file extension
	fileExt := file.GetClientOriginalExtension()

	//filePath
	config := facades.Config()
	basePath := config.Env("APP_URL")

	doctype := req.Doctype

	//validasi folder
	if (doctype == "logo" || doctype == "images" || doctype == "documents" || doctype == "profiles") != true {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{
			"status":  http.StatusBadRequest,
			"success": false,
			"message": lang.MsgDoctypeError,
			"doctype": req.Doctype,
		})
	}

	filePath := fmt.Sprintf("%s/%s/%s/%s", basePath, "api/v1/media/show-file", req.Doctype, fileName)

	//size validation
	fileSize, _ := file.Size()

	if fileSize > 25000000 {
		return ctx.Response().Status(http.StatusRequestEntityTooLarge).Json(http.Json{
			"status":  http.StatusRequestEntityTooLarge,
			"success": false,
			"message": lang.MsgFileToLarge,
		})
	}

	//store proses
	file.StoreAs("./"+req.Doctype, fileName)

	//save to file management table

	_, errsvc := c.FileManagementService.Store(req.Doctype, fileName, fileExt, fileSize, fileType, 0, "")

	if errsvc != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": lang.MsgError,
		})
	}

	return ctx.Response().Status(200).Json(http.Json{
		"status":  200,
		"success": true,
		"message": lang.MsgUploadSuccess,
		"name":    fileName,
		"size":    fileSize,
		"type":    fileType,
		"extn":    fileExt,
		"path":    filePath,
	})
}

func (r MediaController) ShowFile(ctx http.Context) http.Response {
	filename := ctx.Request().Input("filename")
	doctype := ctx.Request().Input("doctype")

	return ctx.Response().File("./storage/app/" + doctype + "/" + filename)
}
