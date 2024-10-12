package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
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

type UserDeleteResponse struct {
	Uuid string `json:"id"`
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

func (s *UserService) Store(req requests.UserRequest) (UserResponse, error) {
	var user models.User

	seederPassword := helpers.GeneratePassword(8)
	hashedpassword, _ := helpers.HashPassword(seederPassword)

	user.Uuid = uuid.New().String()
	user.Name = req.Name
	user.Email = req.Email
	user.Password = hashedpassword
	user.Authent = req.Authent
	user.Phone = req.Phone
	user.Status = req.Status

	err := facades.Orm().Query().Save(&user)

	var tempStatus StatusResponse

	if user.Status == true {
		tempStatus = StatusResponse{
			Color: "green",
			Text:  "Aktif",
		}
	} else {
		tempStatus = StatusResponse{
			Color: "red",
			Text:  "Tidak Aktif",
		}
	}

	if err != nil {
		return UserResponse{}, err
	}

	userResponse := UserResponse{
		Uuid:  user.Uuid,
		Name:  strings.ToUpper(user.Name),
		Email: user.Email,

		Authent: strings.ToUpper(user.Authent),
		Status:  tempStatus,
	}

	//send message
	config := facades.Config()
	wa_title := helpers.AnyToString(config.Env("WA_TITLE"))
	wa_desc := helpers.AnyToString(config.Env("WA_DESC"))
	wa_company := helpers.AnyToString(config.Env("WA_COMPANY"))
	wa_slogan := helpers.AnyToString(config.Env("WA_SLOGAN"))

	message := "*" + wa_title + "* \r\n"
	message = message + wa_desc
	message = message + "\r\n\r\nHalo... \r\n"
	message = message + strings.ToUpper(user.Name)
	message = message + "\r\n\r\nSelamat Anda telah terdaftar sebagai akun pengguna pada sistem kami dengan data akun sebagai berikut :"
	message = message + "\r\nNama pengguna :  " + user.Email
	message = message + "\r\nKata Sandi :  " + seederPassword
	message = message + " \r\n\r\n"
	message = message + wa_slogan
	message = message + "\r\n" + wa_company

	_, errorwa := helpers.SendMessage(req.Phone, message)

	if errorwa != nil {
		fmt.Println(errorwa)
	}

	return userResponse, nil
}

func (s *UserService) Update(req requests.UserRequest, id string) (UserResponse, error) {
	var model models.User

	err := facades.Orm().Query().Where("uuid", id).First(&model)

	if err != nil {
		return UserResponse{}, err
	}

	model.Name = req.Name
	model.Phone = req.Phone
	model.Status = req.Status

	err = facades.Orm().Query().Save(&model)

	if err != nil {
		return UserResponse{}, err
	}

	var tempStatus StatusResponse

	if model.Status == true {
		tempStatus = StatusResponse{
			Color: "green",
			Text:  "Aktif",
		}
	} else {
		tempStatus = StatusResponse{
			Color: "red",
			Text:  "Tidak Aktif",
		}
	}

	userResponse := UserResponse{
		Uuid:    model.Uuid,
		Name:    strings.ToUpper(model.Name),
		Email:   model.Email,
		Authent: strings.ToUpper(model.Authent),
		Status:  tempStatus,
	}

	return userResponse, nil
}

func (s *UserService) Delete(id string) (UserDeleteResponse, error) {

	_, err := facades.Orm().Query().Where("uuid = ?", id).Delete(&models.User{})

	if err != nil {
		return UserDeleteResponse{}, err
	}

	result := UserDeleteResponse{
		Uuid: id,
	}

	return result, nil

}

func (s *UserService) UpdateProfile(ctx http.Context, req requests.UserRequest) (bool, error) {
	var model models.User
	var userauth models.User

	err := facades.Auth(ctx).User(&userauth)

	if err != nil {
		return false, err
	}

	err = facades.Orm().Query().Where("uuid", userauth.Uuid).First(&model)

	if err != nil {
		return false, err
	}

	model.Name = req.Name
	model.Avatar = req.Avatar

	err = facades.Orm().Query().Save(&model)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *UserService) CheckUser(email string) (bool, error) {
	var user models.User
	var count int64

	err := facades.Orm().Query().Model(user).Where("email", email).Count(&count)

	if err != nil {
		return false, err
	}

	if count > 1 {
		return true, nil
	}

	return false, nil
}
