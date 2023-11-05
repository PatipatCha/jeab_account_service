package services

import (
	"log"
	"math/rand"
	"time"

	"github.com/PatipatCha/jeab_ta_service/app/databases"
	"github.com/PatipatCha/jeab_ta_service/app/model"
)

func CheckUserId(mobile_number string) (model.UsersEntity, error) {
	var entity model.UsersEntity
	db, err := databases.ConnectAccountDB()
	if err != nil {
		log.Fatal(err)
		return entity, err
	}

	result := db.Table("users").Where("mobile = ?", mobile_number).Where("status = ?", "active").Scan(&entity)
	if result.RowsAffected <= 0 {
		return entity, err
	}

	return entity, nil
}

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
