package services

import (
	"math/rand"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/model"
)

func SendOTPServer(mobile_number string) (model.MobileOTPSignInResponse, error) {
	var res model.MobileOTPSignInResponse
	verifyRef := randomString(6)
	res = model.MobileOTPSignInResponse{
		VerifyRef: verifyRef,
		Status:    "INPROCESS",
	}
	return res, nil
}

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
