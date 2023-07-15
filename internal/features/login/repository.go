package login

import (
	"jennifer/dealls-tech-test/internal/domain/models"

	"gorm.io/gorm"
)

func FindUserByUsername(db *gorm.DB, username string) (models.Users, error) {
	var user models.Users
	result := db.Where(&models.Users{Username: username}).First(&user)
	return user, result.Error
}
