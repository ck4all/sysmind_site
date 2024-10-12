package services

import (
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
	"strings"
)

type CategoryService struct {
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

type FetchCategoryResponse struct {
	Uuid     string `json:"id"`
	Name     string `json:"name"`
	SlugName string `json:"slug_name"`
	Urutan   int8   `json:"urutan"`
}

func (c *CategoryService) Fetch(page int, itemsPerPage int, keyWord string, sortBy string, sortMode string) ([]FetchCategoryResponse, helpers.MetaResponse, error) {
	var model []models.Category

	var total int64

	query := facades.Orm().Query()

	if sortBy != "" && sortMode != "" {
		query = query.Order(sortBy + " " + sortMode)
	}

	if keyWord != "" {
		query = query.Where("name like ?", keyWord+"%")
	}

	err := query.Paginate(page, itemsPerPage, &model, &total)

	if err != nil {
		return nil, helpers.MetaResponse{}, err
	}

	var datas []FetchCategoryResponse

	for _, element := range model {
		datas = append(datas, FetchCategoryResponse{
			Uuid:     element.Uuid,
			Name:     strings.ToUpper(element.Name),
			SlugName: element.SlugName,
			Urutan:   element.Urutan,
		})
	}

	meta := helpers.SetMeta(page, itemsPerPage, total)

	return datas, meta, err

}
