package server

import (
	"context"
	"fmt"

	"jennifer/dealls-tech-test/internal/server/config"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Name string
}

var db *gorm.DB

func Conn(ctx context.Context) (*gorm.DB, error) {
	if db == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			viper.GetString(config.DatabaseHost),
			viper.GetString(config.DatabaseUser),
			viper.GetString(config.DatabasePassword),
			viper.GetString(config.DatabaseName),
			viper.GetString(config.DatabasePort),
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn))

		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
