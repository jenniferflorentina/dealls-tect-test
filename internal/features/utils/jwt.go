package utils

import (
	"errors"
	"net/http"

	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/server/constants"
)

func GetAuthClaims(r *http.Request) (*others.AuthClaims, error) {
	ctx := r.Context()
	claims := ctx.Value(constants.AuthClaimsCtx)

	if claims == nil {
		return nil, errors.New("auth middleware is not applied to current endpoint")
	}

	return claims.(*others.AuthClaims), nil
}
