package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/models"
	"strings"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInfoResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Authent    string `json:"authent"`
	AvatarPath string `json:"avatar_path"`
}

/**
Register Function
*/

func (a *AuthService) Register(req requests.AuthRegisterRequest) (RegisterResponse, error) {
	var user = models.User{}

	seedPassword := helpers.GeneratePassword(8)
	hashedPassword, _ := helpers.HashPassword(seedPassword)

	user.Uuid = uuid.New().String()
	user.Name = req.Name
	user.Email = req.Email
	user.Password = hashedPassword
	user.Authent = "user"
	user.Avatar = "/images/user.png"
	user.Phone = req.Phone
	user.Status = true

	err := facades.Orm().Query().Save(&user)

	if err != nil {
		return RegisterResponse{}, err
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
	message = message + "\r\nKata Sandi :  " + seedPassword
	message = message + " \r\n\r\n"
	message = message + wa_slogan
	message = message + "\r\n" + wa_company

	_, errorwa := helpers.SendMessage(req.Phone, message)

	if errorwa != nil {
		fmt.Println(errorwa)
	}

	registerResponse := RegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return registerResponse, nil
}

/*
*
Login Function
*/
func (s *AuthService) Login(req requests.AuthLoginRequest) (models.User, bool) {
	var user models.User

	err := facades.Orm().Query().Where("email", req.Email).Where("status", 1).First(&user)

	if err != nil {
		return user, false
	}

	check, errcheck := helpers.CehckPasswordHash(req.Password, user.Password)

	if errcheck != nil {
		return user, false
	}

	if !check {
		return user, false
	}

	return user, true
}

/*
*
Reset Password
*/
func (a *AuthService) ResetPassword(req requests.AuthResetPasswordRequest) (bool, error) {
	var user models.User

	err := facades.Orm().Query().Where("email", req.Email).First(&user)

	if err != nil {
		return false, err
	}

	seedpassword := helpers.GeneratePassword(8)
	hashpassword, _ := helpers.HashPassword(seedpassword)

	user.Password = hashpassword
	errupdate := facades.Orm().Query().Save(&user)

	if errupdate != nil {
		return false, errupdate
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
	message = message + "\r\n\r\nSelamat Anda telah berhasil  melakukan reset kata sandi, dengan detail sebagai berikut :"
	message = message + "\r\nNama pengguna :  " + user.Email
	message = message + "\r\nKata Sandi :  " + seedpassword
	message = message + " \r\n\r\n"
	message = message + wa_slogan
	message = message + "\r\n" + wa_company

	_, errorwa := helpers.SendMessage(user.Phone, message)

	if errorwa != nil {
		fmt.Println(errorwa)
	}

	return true, nil
}

func (s *AuthService) ChangPassword(ctx http.Context, req requests.AuthChangePasswordRequest) (bool, error) {
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

	hashedpassword, _ := helpers.HashPassword(req.Password)

	model.Password = hashedpassword

	err = facades.Orm().Query().Save(&model)

	if err != nil {
		return false, err
	}

	return true, nil

}

/*
*
Chek Email Function
*/
func (a *AuthService) CheckEmail(email string) bool {
	var user models.User
	var count int64

	err := facades.Orm().Query().Model(user).Where("email", email).Count(&count)

	if err != nil {
		return false
	}

	if count < 1 {
		return true
	}

	return false
}

/*
*
Get User Information
*/
func (a *AuthService) UserInfo(ctx http.Context) (UserInfoResponse, error) {
	var model models.User

	err := facades.Auth(ctx).User(&model)

	if err != nil {
		return UserInfoResponse{}, err
	}

	//find avatar path
	config := facades.Config()
	baseUrl := config.Env("APP_URL")

	filePath := ""
	if model.Avatar != "" {
		filePath = fmt.Sprintf("%s/%s/%s", baseUrl, "api/v1/media/show-file/profiles", model.Avatar)
	} else {
		filePath = "/auth/images/user.png"
	}

	result := UserInfoResponse{
		Name:       model.Name,
		Email:      model.Email,
		Authent:    model.Authent,
		AvatarPath: filePath,
	}

	return result, nil
}

func (a *AuthService) Logout(ctx http.Context) (bool, error) {
	err := facades.Auth(ctx).Logout()

	if err != nil {
		return false, err
	}

	return true, nil
}
