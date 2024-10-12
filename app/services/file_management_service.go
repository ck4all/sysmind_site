package services

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
	"math"
)

type FileManagementService struct {
}

func NewFileManagementService() *FileManagementService {
	return &FileManagementService{}
}

type FetchManajemenFileResponse struct {
	Uuid       string  `json:"uuid"`
	FolderName string  `json:"folder_name"`
	FileName   string  `json:"file_name"`
	Ext        string  `json:"ext"`
	Size       float64 `json:"size"`
	Type       string  `json:"type"`
}

func (f *FileManagementService) Fetch(page int, itemsPerPage int, keyWord string, sortBy string, sortMode string) ([]FetchManajemenFileResponse, helpers.MetaResponse, error) {
	var model []models.FileManagement
	var total int64

	query := facades.Orm().Query()

	if sortBy != "" && sortMode != "" {
		query = query.Order(sortBy + " " + sortMode)
	}

	if keyWord != "" {
		query = query.Where("file_name like ?", keyWord+"%")
	}

	err := query.Paginate(page, itemsPerPage, &model, &total)

	var datas []FetchManajemenFileResponse

	for _, element := range model {
		datas = append(datas, FetchManajemenFileResponse{
			Uuid:       element.Uuid,
			FolderName: element.FolderName,
			FileName:   element.FileName,
			Ext:        element.Ext,
			Size:       helpers.ByteToKiloByte(element.Size),
		})
	}

	meta := helpers.MetaResponse{
		Total:       total,
		PerPage:     itemsPerPage,
		CurrentPage: page,
		FirstPage:   page,
		LastPage:    math.Ceil(float64(total/int64(itemsPerPage)) + 1),
	}

	return datas, meta, err

}

func (f *FileManagementService) Store(FolderName string, FileName string, Ext string, Size int64, Type string, Used int, UserUuid string) (models.FileManagement, error) {
	var model models.FileManagement

	model.Uuid = uuid.New().String()
	model.FolderName = FolderName
	model.FileName = FileName
	model.Ext = Ext
	model.Size = Size
	model.Type = Type
	model.Used = Used
	model.UserUuid = UserUuid

	err := facades.Orm().Query().Save(&model)

	if err != nil {
		return models.FileManagement{}, err
	}

	return model, nil
}
