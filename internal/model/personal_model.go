package model

import (
	"gorm.io/gorm"
)

type Personal struct {
	gorm.Model
	IDCard    string `json:"id_card" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
}

func (Personal) TableName() string {
	return "personal"
}
