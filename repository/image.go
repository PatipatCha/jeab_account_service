package repository

type Image struct {
	AccountID int    `db:"user_id"`
	ImageID   string `db:"image_id"`
	ImageUrl  string `db:"image_url"`
}

//go:generate mockgen -destination=../mock/mock_repository/mock_account_repository.go bank/repository AccountRepository
type ImageRepository interface {
	GetImage(int) (*Image, error)
	Upload(Image) (*Image, error)
	Delete(int) (*Image, error)
}
