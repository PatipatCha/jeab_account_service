package services

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/PatipatCha/jeab_account_service/app/databases"
	"github.com/PatipatCha/jeab_account_service/app/model"
)

func CheckUserJMasterPassword(usename string, password string) (model.WebUserProfileResponse, error) {
	var user model.WebUserProfileResponse

	db, err := databases.ConnectMasterDB()
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	userId := strings.ToUpper(usename)

	txtBytes := []byte(password)
	passCode := base64.StdEncoding.EncodeToString(txtBytes)

	db.Table("master_user").
		// Joins("LEFT JOIN operation_center_password on operation_center_password.jeab_id = operation_center_user.jeab_id").
		// Where("operation_center_user.username = ? AND operation_center_password.password = ?", userId, passCode).
		Where("username = ? AND password = ?", userId, passCode).
		Where("role = ?", "master").
		Where("status = ?", "active").
		Scan(&user)

	user.Firstname = "Master"
	user.Surname = "Master"
	user.ImageUrl = ""

	return user, nil
}
