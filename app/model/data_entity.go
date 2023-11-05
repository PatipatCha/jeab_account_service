package model

import (
	"gorm.io/gorm"
)

type UsersEntity struct {
	gorm.Model
	UserId string `db:"user_id" json:"user_id"`
	Mobile string `db:"mobile" json:"mobile"`
	Role   string `db:"role" json:"role"`
	Status string `db:"status" json:"status"`
}

type AddressEntity struct {
}

type ProfileEntity struct {
}
