package login

import (
	"net/http"

	"jennifer/dealls-tech-test/internal/domain/models"
	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/features/utils"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func UseCase(db *gorm.DB, username string, password string) (models.Users, string, *others.BaseResponse) {
	var errorResponse others.BaseResponse
	user, err := FindUserByUsername(db, username)
	if err != nil {
		log.Error().Err(err).Msg("user not found")
		errorResponse = others.BaseResponse{
			Message:    "Username or Password Wrong",
			StatusCode: http.StatusNotFound,
		}
		return models.Users{}, "", &errorResponse
	}

	if !utils.ComparePasswords(user.Password, []byte(password)) {
		log.Error().Err(err).Msg("wrong password")
		errorResponse = others.BaseResponse{
			Message:    "Username or Password Wrong",
			StatusCode: http.StatusNotFound,
		}
		return models.Users{}, "", &errorResponse
	}

	jwtTokenReq := others.Claims{
		UserID:   user.ID,
		Username: user.Username,
	}
	token, err := utils.GenerateJwtToken(jwtTokenReq)
	if err != nil {
		log.Error().Err(err).Msg("Error Generating Token")
		errorResponse = others.BaseResponse{
			Message:    "Error Generating Token",
			StatusCode: http.StatusInternalServerError,
		}
		return models.Users{}, "", &errorResponse
	}

	return user, token, nil
}
