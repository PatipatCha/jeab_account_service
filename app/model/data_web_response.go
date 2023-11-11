package model

type WebUserProfileResponse struct {
	Firstname string `db:"firstname" json:"firstname"`
	Surname   string `db:"surname" json:"surname"`
	ImageUrl  string `db:"image_url" json:"image_url"`
	Mobile    string `db:"mobile" json:"mobile"`
	Role      string `db:"role" json:"role"`
}
