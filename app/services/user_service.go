package services

import (
	"encoding/base64"
	"fmt"
	"time"
)

// func SendOTPServer(mobile_number string) (model.MobileOTPSignInResponse, error) {
// 	var res model.MobileOTPSignInResponse
// 	verifyRef := randomString(6)
// 	res = model.MobileOTPSignInResponse{
// 		VerifyRef: verifyRef,
// 		Status:    "INPROCESS",
// 	}
// 	return res, nil
// }

// func randomString(length int) string {
// 	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	rand.Seed(time.Now().UnixNano())

// 	result := make([]byte, length)
// 	for i := range result {
// 		result[i] = charset[rand.Intn(len(charset))]
// 	}

// 	return string(result)
// }

func GenUserId(role string) string {
	// รับเวลาปัจจุบันในรูปแบบ Unix timestamp แบบ nanoseconds
	currentTimeNano := time.Now().UnixNano()

	// แปลงเวลาเป็นไบต์
	timeBytes := []byte(fmt.Sprintf("%d", currentTimeNano))

	// ใช้ base64 เพื่อเข้ารหัสเวลา
	encodedTime := base64.StdEncoding.EncodeToString(timeBytes)

	// ตัดส่วนที่เกินออกเพื่อให้ได้ขนาดที่ต้องการ (6 ตัวอักษร)
	userIDPart := encodedTime[:6]

	// รวม User ID กับ role
	userID := fmt.Sprintf("%s-%s", userIDPart, role)

	fmt.Println("Generated User ID:", userID)

	return userID
}
