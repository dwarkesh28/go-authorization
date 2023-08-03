package initializers

import "github.com/dwarkesh28/g0-authorization-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
