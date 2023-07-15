package middlewares

import (
	"net/http"
	"strings"

	"jennifer/dealls-tech-test/internal/server/config"

	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func CORSHandler(handler http.Handler) http.Handler {
	allowedOrigins := viper.GetString(config.CorsAllowedOrigins)
	allowedMethods := viper.GetString(config.CorsAllowedMethods)

	corsSetting := cors.New(cors.Options{
		AllowedOrigins: strings.Split(allowedOrigins, ","),
		AllowedMethods: strings.Split(allowedMethods, ","),
		AllowedHeaders: []string{"*"},
	})

	return corsSetting.Handler(handler)
}
