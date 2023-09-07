package repository

type Account struct {
	AccountID int    `db:"user_id"`
	IsActive  bool   `db:"is_active"`
	IDCard    string `db:"idcard"`
	FirstName string `db:"first_name"`
	Surname   string `db:"sur_name"`
	Phone     string `db:"phone"`
}

//go:generate mockgen -destination=../mock/mock_repository/mock_account_repository.go bank/repository AccountRepository
type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
}
