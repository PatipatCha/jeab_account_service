package model

import (
	"gorm.io/gorm"
)

type UsersEntity struct {
	gorm.Model
	UserId       string `db:"user_id" json:"user_id"`
	Mobile       string `db:"mobile" json:"mobile"`
	Role         string `db:"role" json:"role"`
	Status       string `db:"status" json:"status"`
	PersonalPDPA string `db:"personal_pdpa" json:"personal_pdpa"`
}

type ProfileEntity struct {
	gorm.Model
	UserId    string `db:"user_id" json:"user_id"`
	Firstname string `db:"firstname" json:"firstname"`
	Surname   string `db:"surname" json:"surname"`
	ImageUrl  string `db:"image_url" json:"image_url"`
	Mobile    string `db:"mobile" json:"mobile"`
	CreatedBy string `db:"created_by" json:"created_by"`
}

type PDPAEntity struct {
	gorm.Model
	UserId             string `db:"user_id" json:"user_id"`
	PersonalPDPA       string `db:"personal_pdpa" json:"personal_pdpa"`
	PersonalExpireDate string `db:"personal_expire_date" json:"personal_expire_date"`
	CreatedBy          string `db:"created_by" json:"created_by"`
}

type UserProfileEntity struct {
	UserId             string `db:"user_id" json:"user_id"`
	Firstname          string `db:"firstname" json:"firstname"`
	Surname            string `db:"surname" json:"surname"`
	ImageUrl           string `db:"image_url" json:"image_url"`
	Mobile             string `db:"mobile" json:"mobile"`
	Role               string `db:"role" json:"role"`
	PersonalPDPA       string `db:"personal_pdpa" json:"personal_pdpa"`
	PersonalExpireDate string `db:"personal_expire_date" json:"personal_expire_date"`
}
