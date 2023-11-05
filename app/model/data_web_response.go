package model

type WebSignInResponse struct {
	Mobile  string               `json:"mobile"`
	Data    WebOTPSignInResponse `json:"data"`
	Message string               `json:"message"`
}

type WebOTPSignInResponse struct {
	VerifyRef string `json:"verify_ref"`
	Status    string `json:"status"`
}
