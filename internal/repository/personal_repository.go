package repository

import (
	"gorm.io/gorm"
)

type PersonalRepository struct {
	gorm.Model
	IDCard    string `db:"id_card" binding:"required"`
	Firstname string `db:"firstname" binding:"required"`
	Surname   string `db:"surname" binding:"required"`
	Nickname  string `db:"nickname" binding:"required"`
	Email     string `db:"email" binding:"required"`
	Mobile    string `db:"mobile" binding:"required"`
}

func (PersonalRepository) TableName() string {
	return "personal"
}
