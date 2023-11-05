package services

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
