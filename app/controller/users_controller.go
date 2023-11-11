package controller

import (
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	// check mobile value
	mobileNumber := c.Query("mobile")
	if mobileNumber == "" {
		output := fiber.Map{
			"mobile":  mobileNumber,
			"data":    fiber.Map{},
			"message": os.Getenv("MOBILE_IS_NULL"),
		}
		return c.JSON(output)
	}

	// check mobile number
	users, _ := services.GetUser(mobileNumber, "", "")
	if users.UserId == "" {
		output := fiber.Map{
			"mobile":  mobileNumber,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_MOBILE_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	//sendOTP
	data, _ := services.SendOTPService(mobileNumber)

	// res to mobile, web
	res := fiber.Map{
		"mobile":  mobileNumber,
		"data":    data,
		"message": os.Getenv("CHECK_MOBILE_INPROCESS"),
	}

	return c.JSON(res)
}

func VaildOTP(c *fiber.Ctx) error {
	var rawBody = model.MobileOTPSignInRequest{}
	var res = fiber.Map{}

	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	mobileNumber := rawBody.Mobile

	// check mobile number
	users, _ := services.GetUser(mobileNumber, "", "")
	if users.UserId == "" {
		output := fiber.Map{
			"mobile":  mobileNumber,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_MOBILE_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	rawBody.Phone = mobileNumber

	// vaildate otp
	vaild, msg, _ := services.VaildateOTPService(rawBody)
	if vaild {
		profileData, _ := services.GetUser(mobileNumber, "", "")
		res = fiber.Map{
			"mobile":  mobileNumber,
			"data":    profileData,
			"message": msg,
		}
	} else {
		res = fiber.Map{
			"mobile":  mobileNumber,
			"data":    fiber.Map{},
			"message": msg,
		}
	}

	return c.JSON(res)
}

func UpdatePDPA(c *fiber.Ctx) error {
	var res = fiber.Map{}
	var rawBody = model.PDPARequest{}
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	// check user ID
	userId := rawBody.UserId
	user, _ := services.GetUser("", userId, "")
	if user.UserId == "" {
		output := fiber.Map{
			"user_id": userId,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_USER_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	msg, _ := services.UpdatePDPA(user.PersonalPDPA, rawBody)
	res = fiber.Map{
		"user_id": userId,
		"data":    fiber.Map{},
		"message": msg,
	}

	return c.JSON(res)
}

func SignInForJMaster(c *fiber.Ctx) error {
	// var res = fiber.Map{}

	var rawBody = model.WebSignInRequest{}
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	userId := rawBody.UserId
	passcode := rawBody.Passcode

	isUser := services.FindUser(userId)
	if !isUser {
		output := fiber.Map{
			"user_id": userId,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_USER_NOT_FOUND"),
		}
		return c.JSON(output)
	}

	// check user passcode
	user, _ := services.CheckUserPasscode(userId, passcode)
	if user.Firstname == "" {
		output := fiber.Map{
			"user_id": userId,
			"data":    fiber.Map{},
			"message": os.Getenv("SIGN_IN_WRONG"),
		}
		return c.JSON(output)
	}

	res := fiber.Map{
		"user_id": userId,
		"data":    user,
		"message": "SIGN IN SUCCESS",
	}

	return c.JSON(res)
}

func GenId(c *fiber.Ctx) error {
	res := services.GenUserId()
	return c.JSON(res)
}

// func GetPDPA(c *fiber.Ctx) error {
// 	userId := c.Query("user_id")
// 	if userId == "" {
// 		output := fiber.Map{
// 			"userId":  userId,
// 			"data":    fiber.Map{},
// 			"message": os.Getenv("USER_ID_IS_NULL"),
// 		}
// 		return c.JSON(output)
// 	}

// 	user, _ := services.GetUser("", userId, "")
// 	res = fiber.Map{
// 		"user_id": userId,
// 		"data": fiber.Map{
// 			"personal_pdpa":        user.PersonalPDPA,
// 			"personal_expire_date": user.PersonalExpireDate,
// 		},
// 		"message": "PDPA Lists",
// 	}

// 	return c.JSON(res)
// }
