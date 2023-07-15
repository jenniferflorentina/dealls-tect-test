package others

import "github.com/google/uuid"

type Claims struct {
	UserID   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
}
