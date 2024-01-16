package model

type WebUserProfileResponse struct {
	JeabID    string `db:"jeab_id" json:"jeab_id"`
	Firstname string `db:"firstname" json:"firstname"`
	Surname   string `db:"surname" json:"surname"`
	ImageUrl  string `db:"image_url" json:"image_url"`
	Mobile    string `db:"mobile" json:"mobile"`
	Role      string `db:"role" json:"role"`
}
