package utils

import (
	"strconv"
	"time"

	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/server/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GenerateJwtToken(claims others.Claims) (string, error) {
	rawSecret := viper.GetString(config.JWTSecretKey)
	jwtSecret := []byte(rawSecret)
	expiredMinute, _ := strconv.Atoi(config.JWTSecretKey)
	expirationDate := time.Now().Add(time.Duration(expiredMinute) * time.Minute).UTC()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, others.AuthClaims{
		Username: claims.Username,
		UserID:   claims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationDate),
		},
	})

	return token.SignedString(jwtSecret)
}
