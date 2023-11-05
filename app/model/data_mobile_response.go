package model

type MobileSignInResponse struct {
	Mobile  string                  `json:"mobile"`
	Data    MobileOTPSignInResponse `json:"data"`
	Message string                  `json:"message"`
}

type MobileOTPSignInResponse struct {
	VerifyRef string `json:"verify_ref"`
	Status    string `json:"status"`
}
