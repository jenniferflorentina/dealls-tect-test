package utils

import (
	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/server/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func ParseToken(token string) (*others.AuthClaims, error) {
	jwtSecret := []byte(viper.GetString(config.JWTSecretKey))
	claims := &others.AuthClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	return claims, err
}
