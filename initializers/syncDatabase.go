package initializers

import "go-authorization-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
