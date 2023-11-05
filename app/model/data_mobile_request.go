package model

type MobileSignInRequest struct {
	Mobile string `db:"mobile" json:"mobile"`
}

type MobileOTPSignInRequest struct {
	Mobile    string `db:"mobile" json:"mobile"`
	OTP       string `db:"otp" json:"otp"`
	VerifyRef string `json:"verify_ref"`
}
