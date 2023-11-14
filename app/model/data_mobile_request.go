package model

type MobileSignInRequest struct {
	Mobile string `json:"mobile"`
}

type ChangeMobileRequest struct {
	NewMobile string `json:"new_mobile"`
}

type MobileOTPRequest struct {
	ChangeMobileRequest
	MobileSignInRequest
	MobileOTPSignInResponse
	OTP   string `json:"otp_code"`
	Phone string `json:"phone"`
}

type PDPARequest struct {
	UserId       string `json:"user_id"`
	PersonalPDPA string `json:"personal_pdpa"`
}

type UserProfileRequest struct {
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
}
