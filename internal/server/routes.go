package server

import (
	"net/http"

	"jennifer/dealls-tech-test/internal/features/register"

	"jennifer/dealls-tech-test/internal/features/login"

	"jennifer/dealls-tech-test/internal/features/healthcheck"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) http.Handler {
	router := httprouter.New()

	router.GET("/health", healthcheck.New())
	router.POST("/authentication/login", login.New(db))
	router.POST("/authentication/register", register.New(db))
	return router
}
