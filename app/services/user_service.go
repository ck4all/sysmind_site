package services

import (
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/models"
	"math"
	"strings"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

type StatusResponse struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

type FetchUserResponse struct {
	Uuid    string         `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Authent string         `json:"authent"`
	Status  StatusResponse `json:"status"`
}

type UserResponse struct {
	Uuid    string         `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Authent string         `json:"authent"`
	Status  StatusResponse `json:"status"`
}

type UserShowResponse struct {
	Uuid    string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Authent string `json:"authent"`
	Phone   string `json:"phone"`
	Status  bool   `json:"status"`
}

/**
Fetch Data Users
*/

func (s *UserService) Fetch(page int, itemsPerPage int, keyWord string, sortBy string, sortMode string) ([]FetchUserResponse, helpers.MetaResponse, error) {
	var users []models.User

	var total int64

	query := facades.Orm().Query()

	query = query.Where("authent != ?", "superadmin")

	if sortBy != "" && sortMode != "" {
		query = query.Order(sortBy + " " + sortMode)
	}

	if keyWord != "" {
		query = query.Where("name like ?", keyWord+"%")
	}

	err := query.Paginate(page, itemsPerPage, &users, &total)

	if err != nil {
		return nil, helpers.MetaResponse{}, err
	}

	var datas []FetchUserResponse

	for _, element := range users {
		var tempstatus StatusResponse
		if element.Status == true {
			tempstatus = StatusResponse{
				Color: "green",
				Text:  "Aktif",
			}
		} else {
			tempstatus = StatusResponse{
				Color: "red",
				Text:  "Tidak Aktif",
			}
		}
		datas = append(datas, FetchUserResponse{
			Uuid:    element.Uuid,
			Name:    strings.ToUpper(element.Name),
			Email:   element.Email,
			Authent: strings.ToUpper(element.Authent),
			Status:  tempstatus,
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

func (s *UserService) Show(id string) (UserShowResponse, error) {
	var model models.User

	err := facades.Orm().Query().Where("uuid", id).First(&model)

	if err != nil {
		return UserShowResponse{}, err
	}

	userShowResponse := UserShowResponse{
		Uuid:    model.Uuid,
		Name:    model.Name,
		Email:   model.Email,
		Authent: model.Authent,
		Phone:   model.Phone,
		Status:  model.Status,
	}

	return userShowResponse, nil
}
