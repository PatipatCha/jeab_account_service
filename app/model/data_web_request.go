package model

type WebSignInRequest struct {
	UserId   string `json:"user_id"`
	Passcode string `json:"passcode"`
}

type JMasterLogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
