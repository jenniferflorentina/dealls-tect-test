package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"jennifer/dealls-tech-test/internal/features/utils"

	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/server/config"
	"jennifer/dealls-tech-test/internal/server/constants"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

type authError struct {
	Auth string `json:"auth"`
}

// Extracts the token payload/claims from the request header,
// it's too much if I put this inside the middleware itself so I separated it.
//
// Note that this function works like any function that has an error,
// basically the `claims` won't be a nil and the `error` can be nil.
func extractTokenClaims(header string) (*others.AuthClaims, error) {
	jwtSecret := []byte(viper.GetString(config.JWTSecretKey))
	splitHeader := strings.Split(header, "Bearer ")
	claims := &others.AuthClaims{}

	if len(splitHeader) == 1 {
		return claims, errors.New("missing bearer authorization")
	}

	rawToken := splitHeader[1]
	_, err := jwt.ParseWithClaims(rawToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	return claims, err
}

func AuthHandler(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		header := r.Header.Get("authorization")
		claims, err := extractTokenClaims(header)

		if err != nil {
			response := others.BaseResponse{
				Message:    err.Error(),
				StatusCode: http.StatusUnauthorized,
				Data:       authError{err.Error()},
			}
			utils.SendResponse(w, response)
			return
		}

		newCtx := context.WithValue(r.Context(), constants.AuthClaimsCtx, claims)
		next(w, r.WithContext(newCtx), p)
	})
}
