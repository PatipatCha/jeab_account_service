package model

import "gorm.io/gorm"

type Account struct {
	// AccountID int    `db:"user_id"`
	// IsActive  bool   `db:"is_active"`
	// IDCard    string `db:"idcard"`
	// FirstName string `db:"first_name"`
	// Surname   string `db:"sur_name"`
	// Phone     string `db:"phone"`
	gorm.Model
	IDCard    string `json:"id_card" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
}

//go:generate mockgen -destination=../mock/mock_repository/mock_account_repository.go bank/repository AccountRepository
type AccountRepository interface {
	// Create(Account) (*Account, error)
	// findById(int) ([]Account, error)
	findAll() ([]Account, error)
}
