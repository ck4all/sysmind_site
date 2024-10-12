package helpers

import (
	"bytes"
	"fmt"
	"github.com/goravel/framework/facades"
	"net/http"
	"net/url"
)

type WhatsappRequest struct {
	ApiKey string `json:"apiKey"`
}

func SendMessage(recieveNumber string, message string) (bool, error) {
	config := facades.Config()
	apikey := config.Env("WA_API_KEY")

	formData := url.Values{}
	formData.Set("apiKey", AnyToString(apikey))
	formData.Set("recieveNumber", recieveNumber)
	formData.Set("message", message)

	reqBody := bytes.NewBufferString(formData.Encode())

	result, err := http.Post("https://api.senderwa.com/api/v2/send-message", "application/x-www-form-urlencoded", reqBody)

	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		return false, err
	}

	return true, nil
}
