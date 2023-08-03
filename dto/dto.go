package dto

import (
	"go-authorization-jwt/models"
	"time"
)

type GetData struct {
	ID        uint64    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:nano"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
}

func ToDto(user models.User) *GetData {
	return &GetData{
		ID:        user.ID,
		CreatedAt: time.Unix(0, user.CreatedAt),
		UpdatedAt: time.Unix(0, user.UpdatedAt),
		Email:     user.Email,
		Password:  user.Password,
	}
}
