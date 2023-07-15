package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"varchar(255)"`
	Password string    `gorm:"varchar(255)"`
	Name     string    `gorm:"varchar(255)"`
	Phone    string    `gorm:"varchar(15)"`
	Email    string    `gorm:"varchar(255)"`
}
