package register

import (
	"net/http"

	"jennifer/dealls-tech-test/internal/domain/models"
	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/features/utils"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func UseCase(db *gorm.DB, username string, password string, name string, phone string, email string) (models.Users, *others.BaseResponse) {
	var errorResponse *others.BaseResponse
	_, err := FindUserByUsername(db, username)
	if err == nil {
		log.Error().Msg("Username already exist!")
		errorResponse = &others.BaseResponse{
			Message:    "Error Creating User, username already exist!",
			StatusCode: http.StatusInternalServerError,
		}
		return models.Users{}, errorResponse
	}

	newUser := models.Users{
		ID:       uuid.New(),
		Username: username,
		Password: utils.HashAndSalt([]byte(password)),
		Name:     name,
		Phone:    phone,
		Email:    email,
	}
	user, err := CreateUser(db, newUser)
	if err != nil {
		log.Error().Err(err).Msg("Error Creating User")
		errorResponse = &others.BaseResponse{
			Message:    "Error Creating User",
			StatusCode: http.StatusInternalServerError,
		}
		return models.Users{}, errorResponse
	}
	return user, errorResponse
}
