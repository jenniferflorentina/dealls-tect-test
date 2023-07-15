package others

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthClaims struct {
	UserID   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}
