package config

import (
	"github.com/rotisserie/eris"
	"github.com/spf13/viper"
)

func validate() {
	requiredKeys := []string{
		DatabaseHost,
		DatabasePort,
		DatabaseName,
		DatabaseUser,
		DatabasePassword,
		JWTSecretKey,
		JWTExpirationTime,
		CorsAllowedOrigins,
		CorsAllowedMethods,
	}

	for _, key := range requiredKeys {
		if viper.Get(key) != nil {
			continue
		}

		panic(key + " doesn't have a value in '.env'")
	}
}

func Setup() {
	defer validate()

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(eris.Wrap(err, "error loading config"))
		}
	}
	viper.AutomaticEnv()
}

const (
	DatabaseUser     = "DB_USER"
	DatabasePassword = "DB_PASSWORD"
	DatabaseHost     = "DB_HOST"
	DatabasePort     = "DB_PORT"
	DatabaseName     = "DB_NAME"

	AdminUserName     = "ADMIN_USER_NAME"
	AdminPassword     = "ADMIN_PASSWORD"
	JWTSecretKey      = "JWT_SECRET_KEY"
	JWTExpirationTime = "JWT_EXPIRATION_TIME"

	CorsAllowedOrigins = "CORS_ALLOWED_ORIGINS"
	CorsAllowedMethods = "CORS_ALLOWED_METHODS"
)
