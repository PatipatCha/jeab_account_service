package repository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm.Model
	Username string `db:"text" binding:"required"`
	Mobile   string `db:"mobile" binding:"required"`
}

func (UserRepository) TableName() string {
	return "users"
}
