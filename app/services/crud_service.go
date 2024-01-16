package services

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PatipatCha/jeab_account_service/app/databases"
	"github.com/PatipatCha/jeab_account_service/app/model"
)

func GetUser(mobile_number string, user_id string, role string) (model.UserProfileEntity, error) {
	var entity model.UserProfileEntity
	db, err := databases.ConnectAccountDB()
	if err != nil {
		log.Fatal(err)
		return entity, err
	}

	var raw = ""
	if role != "" {
		raw = "users.role = " + role
	}

	sqlSelect := "users.user_id, profile.firstname, profile.surname, users.mobile, profile.image_url, users.role, pdpa.personal_pdpa, TO_CHAR(pdpa.personal_expire_date, 'YYYY-MM-DD') as personal_expire_date, staff.seoc_id as seoc_id"
	sqlJoinProfle := "LEFT JOIN profile on profile.user_id = users.user_id"
	sqlJoinPdpa := "LEFT JOIN pdpa on pdpa.user_id = users.user_id"

	var result = db.Table("users").Select(sqlSelect).Joins(sqlJoinProfle).Joins(sqlJoinPdpa).Joins("LEFT JOIN staff on staff.user_id = users.user_id").Where("users.user_id = ?", user_id).Or("users.mobile = ?", mobile_number).Where("users.status = ?", "active").Where(raw).Scan(&entity)
	fmt.Println(entity)
	if result.RowsAffected <= 0 {
		return entity, err
	}

	return entity, nil
}

func FindUser(user_id string, role string) bool {
	var entity model.UserProfileEntity

	db, err := databases.ConnectAccountDB()
	if err != nil {
		log.Fatal(err)
		return false
	}

	db.Table("users").Where("users.user_id = ?", user_id).Where("users.role = ?", role).First(&entity)

	return entity.UserId != ""
}

func CheckUserPasscode(user_id string, passcode string) (model.WebUserProfileResponse, error) {
	var output model.WebUserProfileResponse

	db, err := databases.ConnectAccountDB()
	if err != nil {
		log.Fatal(err)
		return output, err
	}

	userId := strings.ToUpper(user_id)

	txtBytes := []byte(passcode)
	passCode := base64.StdEncoding.EncodeToString(txtBytes)

	db.Table("users").Select("profile.firstname, profile.surname, users.mobile, profile.image_url, users.role").Joins("LEFT JOIN profile on profile.user_id = users.user_id").Joins("LEFT JOIN passcode on passcode.user_id = users.user_id").Where("users.user_id = ?", userId).Where("passcode.passcode = ?", passCode).Where("users.status = ?", "active").Scan(&output)

	return output, nil
}

// func GetProfile(mobile_number string, user_id string) (model.ProfileEntity, error) {
// 	var entity model.ProfileEntity
// 	db, err := databases.ConnectAccountDB()
// 	if err != nil {
// 		log.Fatal(err)
// 		return entity, err
// 	}

// 	result := db.Table("profile").Select("profile.user_id, profile.firstname, profile.surname, profile.image_url, users_pdpa.personal_pdpa, TO_CHAR(users_pdpa.personal_expire_date, 'YYYY-MM-DD') as personal_expire_date").Joins("LEFT JOIN users_pdpa on users_pdpa.user_id = profile.user_id").Where("profile.user_id = ?", user_id).Or("profile.mobile = ?", mobile_number).First(&entity)
// 	if result.RowsAffected <= 0 {
// 		return entity, err
// 	}

// 	return entity, nil

// }

func UpdatePDPA(existingPDPA string, request model.PDPARequest) (string, error) {
	var res = "INSERT SUCCESS"
	userId := string(request.UserId)
	personalPDPA := string(request.PersonalPDPA)
	personalExpireDate := ConvertTextExpireDate(3)

	entity := model.PDPAEntity{
		UserId:             userId,
		PersonalPDPA:       personalPDPA,
		PersonalExpireDate: personalExpireDate,
		CreatedBy:          userId,
	}

	db, err := databases.ConnectAccountDB()
	if err != nil {
		log.Fatal(err)
		return os.Getenv("ERROR_DB"), err
	}

	if existingPDPA != "" {
		if err := db.Table("pdpa").Where("pdpa.user_id = ?", userId).Updates(entity).Error; err != nil {
			log.Fatal(err)
			return "UPDATE ERROR", err
		}
		res = "UPDATE SUCCESS"
	} else {
		if err := db.Table("pdpa").Create(&entity).Error; err != nil {
			log.Fatal(err)
			return "INSERT ERROR", err
		}
	}

	return res, nil
}

func UpdateProfile(userId string, request model.UserProfileRequest) (string, error) {

	db, err := databases.ConnectAccountDB()
	if err != nil {
		log.Fatal(err)
		return os.Getenv("ERROR_DB"), err
	}

	var entity = model.ProfileEntity{}
	entity.Firstname = request.Firstname
	entity.Surname = request.Surname

	if err := db.Table("profile").Where("user_id = ?", userId).Updates(&entity).Error; err != nil {
		return os.Getenv("UPDATE_PROFILE_ERROR"), nil
	}

	return os.Getenv("UPDATE_PROFILE_SUCCESS"), nil
}

func UpdateMobile(userId string, request model.MobileOTPRequest) (string, error) {
	msg := os.Getenv("UPDATE_PROFILE_SUCCESS")

	db, err := databases.ConnectAccountDB()
	if err != nil {
		log.Fatal(err)
		return os.Getenv("ERROR_DB"), err
	}

	tx := db.Begin()

	if err := tx.Table("users").Where("user_id = ?", userId).Where("mobile = ?", request.Mobile).Update("mobile", request.NewMobile).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
		msg = os.Getenv("UPDATE_USER_ERROR")
	}

	if err := tx.Table("profile").Where("user_id = ?", userId).Where("mobile = ?", request.Mobile).Update("mobile", request.NewMobile).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
		msg = os.Getenv("UPDATE_PROFILE_ERROR")
	}

	tx.Commit()

	return msg, nil
}

// func SaveData(request model.TimeAttendanceCheckInOutRequest) (model.TimeAttendanceEntity, error) {
// 	entity := model.TimeAttendanceEntity{
// 		UserId:        string(request.UserId),
// 		CheckDateTime: string(request.CheckDateTime),
// 		ProjectId:     string(request.ProjectId),
// 		ProjectPlace:  string(request.ProjectPlace),
// 		CheckStatus:   strings.ToLower(request.CheckStatus),
// 		CreatedBy:     request.CreatedBy,
// 		ImageUrl:      request.ImageUrl,
// 		RefId:         request.RefId,
// 	}

// 	db, err := databases.ConnectTADB()
// 	if err != nil {
// 		return entity, err
// 	}

// 	err = db.Table("time_attendance").Create(&entity).Scan(&entity).Error

// 	// err = db.Table("time_attendance").Find("user_id").Error

// 	return entity, err
// }

// func GetReportForMobile(user_id string, month string) (bool, []model.TimeAttendanceReportMobileEntity, string) {
// 	var ta_entity = []model.TimeAttendanceReportMobileEntity{}
// 	var msg = "Record Lists"

// 	db, err := databases.ConnectTADB()
// 	if err != nil {
// 		return false, ta_entity, string(err.Error())
// 	}

// 	//SQL RAW SELECT

// 	var sqlRawWhereMonth = "AND EXTRACT( MONTH FROM a.check_date_time ) = EXTRACT( MONTH FROM LOCALTIMESTAMP AT TIME ZONE 'utc+7' ) "
// 	if month != "" {
// 		sqlRawWhereMonth = "AND EXTRACT( MONTH FROM a.check_date_time ) = ? "
// 	}

// 	sqlRawSelectPlace := "SELECT DATE (a.check_date_time) AS date, a.project_place AS project_place, TO_CHAR(a.check_date_time, 'HH24:MI') AS check_in_time, TO_CHAR(b.check_date_time, 'HH24:MI') AS check_out_time, CASE WHEN EXTRACT( DAY FROM a.check_date_time) != EXTRACT( DAY FROM b.check_date_time) THEN TO_CHAR(b.check_date_time, 'YYYY-MM-DD') ELSE '' END AS check_out_remark "
// 	sqlRawSelctB := ",FLOOR(EXTRACT(EPOCH FROM b.check_date_time::timestamp - a.check_date_time::timestamp)/3600)::int2 AS total_hour, ABS(EXTRACT( MINUTE FROM a.check_date_time) + EXTRACT( MINUTE FROM b.check_date_time)) AS total_minute "
// 	sqlRawFrom := "FROM time_attendance a,time_attendance b "
// 	sqlRawWhereUserId := "WHERE a.user_id = ? "
// 	sqlRawWhere := "AND a.check_status = 'checkin' AND b.check_status = 'checkout' AND a.ref_id = b.ref_id "
// 	sqlRawOrderBy := "ORDER BY a.check_date_time DESC"

// 	var sqlRaw = sqlRawSelectPlace + sqlRawSelctB + sqlRawFrom + sqlRawWhereUserId + sqlRawWhere + sqlRawWhereMonth + sqlRawOrderBy

// 	if month != "" {
// 		db.Raw(sqlRaw, user_id, month).Scan(&ta_entity)
// 	} else {
// 		db.Raw(sqlRaw, user_id).Scan(&ta_entity)
// 	}

// 	return true, ta_entity, msg

// }

// func GetReportForWeb(findUserId string) ([]model.TimeAttendanceDashboardList, string) {

// 	var ta_dashboard = []model.TimeAttendanceDashboardList{}

// 	db, err := databases.ConnectTADB()
// 	if err != nil {
// 		return ta_dashboard, "Database Not Connected"
// 	}

// 	sqlRawA := "SELECT a.user_id AS \"user_id\", a.project_place AS \"project_place\", TO_CHAR( a.check_date_time :: DATE, 'dd-mm-yyyy' ) AS \"check_in_date\", a.image_url AS \"check_in_image\", TO_CHAR(a.check_date_time, 'HH24:MI') AS \"check_in_time\", TO_CHAR( b.check_date_time :: DATE, 'dd-mm-yyyy' ) AS \"check_out_date\", TO_CHAR(b.check_date_time, 'HH24:MI') AS \"check_out_time\", b.image_url AS \"check_out_image\" FROM time_attendance a, time_attendance b "
// 	sqlRawB := "WHERE a.check_status = 'checkin' AND b.check_status = 'checkout' AND a.ref_id = b.ref_id "
// 	_ = "AND EXTRACT( MONTH FROM a.check_date_time ) = EXTRACT( MONTH FROM LOCALTIMESTAMP AT TIME ZONE 'utc+7' ) "
// 	sqlRawC := "AND a.user_id = ? "
// 	sqlRawD := "ORDER BY a.check_date_time DESC"

// 	var sqlRaw = sqlRawA + sqlRawB + sqlRawD
// 	if findUserId != "" {
// 		sqlRaw = sqlRawA + sqlRawB + sqlRawC + sqlRawD
// 		db.Raw(sqlRaw, findUserId).Scan(&ta_dashboard)
// 	} else {
// 		db.Raw(sqlRaw).Scan(&ta_dashboard)
// 	}

// 	return ta_dashboard, "Get Record List"
// }

// func GetReportNow() model.TimeAttendanceReportList {
// 	var ta_report = model.TimeAttendanceReportList{}
// 	content, err := ioutil.ReadFile("./app/json/record_mockup_test_mobile.json")
// 	if err != nil {
// 		log.Fatal("Error when opening file: ", err)
// 	}

// 	err = json.Unmarshal(content, &ta_report)

// 	return ta_report
// }

// func GetReport(c *fiber.Ctx) error {
// 	var ta []model.TimeAttendanceEntity
// 	db, err := databases.ConnectDB()
// 	if err != nil {
// 		return err
// 	}

// 	// userId := c.Params("userId")
// 	res := db.Table("time_attendance").Find(&ta).Error
// 	if err := res; err != nil {
// 		return err
// 	}

// 	println()

// 	return c.JSON(ta)
// }
