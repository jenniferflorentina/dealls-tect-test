package utils

import (
	"net/http"

	"jennifer/dealls-tech-test/internal/domain/others"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func Validator(w http.ResponseWriter, s interface{}) error {
	val := validator.New()
	err := val.Struct(s)

	if err != nil {
		log.Error().Err(err).Msg("Request Payload Validation Not Satisfied")
		SendResponse(w, others.BaseResponse{
			Message:    "Request Payload Validation Not Satisfied",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return err
	}
	return nil
}
