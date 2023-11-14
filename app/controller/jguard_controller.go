package controller

import (
	"os"
	"strings"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func SignInAndChangeMobileHandler(c *fiber.Ctx) error {
	var mobileNumber string
	var rawBody = model.MobileSignInRequest{}

	// 1) check raw body
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	// 2) trim mobile number
	if rawBody.Mobile != "" {
		mobileNumber = strings.TrimSpace(rawBody.Mobile)
	}

	// 3) check user by mobile number
	users, _ := services.GetUser(mobileNumber, "", "")
	if users.UserId == "" {
		res := fiber.Map{
			"mobile":  mobileNumber,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_MOBILE_NOT_FOUND"),
		}
		return c.Status(404).JSON(res)
	}

	// 4) sendOTP
	data, _ := services.SendOTPService(mobileNumber)
	res := fiber.Map{
		"mobile":  mobileNumber,
		"data":    data,
		"message": os.Getenv("CHECK_MOBILE_INPROCESS"),
	}

	return c.Status(200).JSON(res)
}

func VaildOTPHandler(c *fiber.Ctx) error {
	var mobileNumber string
	var rawBody = model.MobileOTPRequest{}
	var res = fiber.Map{}

	// 1) check raw body
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	// 2) trim mobile number
	if rawBody.Mobile != "" {
		mobileNumber = strings.TrimSpace(rawBody.Mobile)
	}

	// 3) check user by mobile number
	users, _ := services.GetUser(mobileNumber, "", "")
	if users.UserId == "" {
		output := fiber.Map{
			"mobile":  mobileNumber,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_MOBILE_NOT_FOUND"),
		}
		return c.Status(404).JSON(output)
	}

	// 4) vaildate otp
	rawBody.Phone = mobileNumber
	vaild, msg, _ := services.VaildateOTPService(rawBody)
	res = fiber.Map{
		"data":    fiber.Map{},
		"message": msg,
	}

	userId := c.Query("user_id")
	if userId != "" {
		res["user_id"] = userId
		// 4.1) by user id
		if vaild {
			msg, _ := services.UpdateMobile(userId, rawBody)
			res["message"] = msg
		}
	} else {
		// 4.2) by mobile number
		res["mobile"] = mobileNumber
		if vaild {
			profileData, _ := services.GetUser(mobileNumber, "", "")
			res["data"] = profileData
		}
	}

	return c.Status(200).JSON(res)
}

func GetUserHandler(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	mobile := c.Query("mobile")
	user, _ := services.GetUser(mobile, userId, "")
	res := fiber.Map{
		"user_id": userId,
		"data":    user,
		"message": os.Getenv("USER_LIST"),
	}
	return c.Status(200).JSON(res)
}

func GenId(c *fiber.Ctx) error {
	res := services.GenUserId("MASTER")
	return c.JSON(res)
}

func GetPDPAHandler(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	user, _ := services.GetUser("", userId, "")
	res := fiber.Map{
		"user_id": userId,
		"data": fiber.Map{
			"personal_pdpa":        user.PersonalPDPA,
			"personal_expire_date": user.PersonalExpireDate,
		},
		"message": "PDPA Lists",
	}
	return c.Status(200).JSON(res)
}

func UpdatePDPAHandler(c *fiber.Ctx) error {
	var userId string
	var res = fiber.Map{}
	var rawBody = model.PDPARequest{}
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	if rawBody.UserId != "" {
		userId = strings.TrimSpace(rawBody.UserId)
	}

	// check user ID
	user, _ := services.GetUser("", userId, "")
	if user.UserId == "" {
		res := fiber.Map{
			"user_id": userId,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_USER_NOT_FOUND"),
		}
		return c.Status(404).JSON(res)
	}

	msg, _ := services.UpdatePDPA(user.PersonalPDPA, rawBody)
	res = fiber.Map{
		"user_id": userId,
		"data":    fiber.Map{},
		"message": msg,
	}

	return c.Status(200).JSON(res)
}

func UpdateProfileHandler(c *fiber.Ctx) error {
	userId := c.Query("user_id")

	var rawBody = model.UserProfileRequest{}
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	if userId == "" {
		res := fiber.Map{
			"user_id": userId,
			"data":    fiber.Map{},
			"message": os.Getenv("USER_ID_IS_NULL"),
		}
		return c.Status(404).JSON(res)
	}

	users, _ := services.GetUser("", userId, "")
	if users.UserId != userId {
		res := fiber.Map{
			"mobile":  userId,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_USER_NOT_FOUND"),
		}
		return c.Status(404).JSON(res)
	}

	msg, _ := services.UpdateProfile(userId, rawBody)
	res := fiber.Map{
		"user_id": userId,
		"data":    fiber.Map{},
		"message": msg,
	}
	return c.Status(200).JSON(res)
}
