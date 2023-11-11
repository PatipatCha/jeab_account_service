package model

type MobileSignInRequest struct {
	Mobile string `json:"mobile"`
}

type MobileOTPSignInRequest struct {
	MobileSignInRequest
	MobileOTPSignInResponse
	OTP   string `json:"otp_code"`
	Phone string `json:"phone"`
}

type PDPARequest struct {
	UserId       string `json:"user_id"`
	PersonalPDPA string `json:"personal_pdpa"`
}
