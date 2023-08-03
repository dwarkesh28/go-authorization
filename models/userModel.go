package models

type User struct {
	ID        uint64 `json:"id" db:"id"`
	CreatedAt int64  `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Email     string `gorm:"unique"`
	Password  string
}
