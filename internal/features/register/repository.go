package register

import (
	"jennifer/dealls-tech-test/internal/domain/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.Users) (models.Users, error) {
	result := db.Create(&user)
	return user, result.Error
}

func FindUserByUsername(db *gorm.DB, username string) (models.Users, error) {
	var user models.Users
	result := db.Where(&models.Users{Username: username}).First(&user)
	return user, result.Error
}
