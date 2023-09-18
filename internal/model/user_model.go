package model

import "time"

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `json:"text" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  time.Time
}

func (User) TableName() string {
	return "users"
}
