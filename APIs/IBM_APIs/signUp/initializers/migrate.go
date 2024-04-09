package initializers

import (
	"signUp/models"
)

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
}
