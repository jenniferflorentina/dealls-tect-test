package login

import (
	"encoding/json"
	"io"
	"net/http"

	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/features/utils"

	"github.com/rs/zerolog/log"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type auth struct {
	db *gorm.DB
}

func New(db *gorm.DB) httprouter.Handle {
	return auth{db: db}.UpdateTokenAuth
}

type Request struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (a auth) UpdateTokenAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Parsing JSON body failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestBody Request
	err = json.Unmarshal(jsonBody, &requestBody)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshaling JSON body failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = utils.Validator(w, requestBody)
	if err != nil {
		return
	}

	user, token, response := UseCase(a.db, requestBody.Username, requestBody.Password)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Login Succeed",
			StatusCode: http.StatusOK,
			Data: Response{
				UserID:   user.ID.String(),
				Name:     user.Name,
				Username: user.Username,
				Token:    token,
			},
		}
	}
	utils.SendResponse(w, *response)
}
