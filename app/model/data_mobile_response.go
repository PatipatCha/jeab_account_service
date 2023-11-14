package model

type MobileOTPSignInResponse struct {
	Token string `json:"token"`
}
type MobileSignInResponse struct {
	Mobile  string                  `json:"mobile"`
	Data    MobileOTPSignInResponse `json:"data"`
	Message string                  `json:"message"`
}

type UserProfileResponse struct {
	UserProfileRequest
}
