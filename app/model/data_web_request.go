package model

type WebSignInRequest struct {
	UserId   string `json:"user_id"`
	Passcode string `json:"passcode"`
}
