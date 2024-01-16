package services

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/PatipatCha/jeab_account_service/app/model"
	"github.com/gofiber/fiber/v2"
)

func SendOTPService(mobileNumber string) (model.MobileOTPSignInResponse, error) {
	var baseUrl = os.Getenv("MESSAGE_SERVICE_URL") + os.Getenv("MESSAGE_SERVICE_SEND_OTP_URL")
	var project_key = os.Getenv("PROJECT_KEY")

	postData := fiber.Map{
		"project_key": project_key,
		"phone":       mobileNumber,
	}

	res := model.MobileOTPSignInResponse{}

	jsonBody, _ := json.Marshal(postData)
	response, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return res, err
	}
	defer response.Body.Close()

	responseBody, _ := ioutil.ReadAll(response.Body)
	var responseData model.OTPResponseData
	json.Unmarshal(responseBody, &responseData)

	// Check the response status
	if response.StatusCode == http.StatusCreated {
		return res, nil
	}

	res = model.MobileOTPSignInResponse{
		Token: responseData.Result.Token,
	}

	return res, nil

}

func VaildateOTPService(postData model.MobileOTPRequest) (bool, string, error) {
	var baseUrl = os.Getenv("MESSAGE_SERVICE_URL") + os.Getenv("MESSAGE_SERVICE_VAILDATE_OTP_URL")
	jsonBody, _ := json.Marshal(postData)
	response, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return false, "", nil
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	var responseData model.OTPVaildateResponseData
	json.Unmarshal(responseBody, &responseData)

	// *Develop Mode
	// var responseData model.OTPVaildateResponseData
	// responseData.Result.Status = true
	// responseData.Code = "0000"
	// var err error
	//

	if responseData.Code == "1006" {
		return false, os.Getenv("VALID_OTP_EXPIRED"), err
	}

	if !responseData.Result.Status {
		return false, os.Getenv("VALID_OTP_FAIL"), err
	}

	return true, os.Getenv("VALID_OTP_SUCCESS"), nil

}
