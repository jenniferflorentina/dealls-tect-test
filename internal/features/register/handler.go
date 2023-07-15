package register

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/features/utils"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

type Request struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=20"`
	Name     string `json:"name" validate:"required,max=255"`
	Phone    string `json:"phone" validate:"required,max=15"`
	Email    string `json:"email" validate:"required"`
}

type Response struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func New(db *gorm.DB) httprouter.Handle {
	return handler{db: db}.Handle
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	user, response := UseCase(h.db, requestBody.Username, requestBody.Password, requestBody.Name, requestBody.Phone, requestBody.Email)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "User Created",
			StatusCode: http.StatusCreated,
			Data: Response{
				ID:       user.ID.String(),
				Username: user.Username,
				Name:     user.Name,
				Phone:    user.Phone,
				Email:    user.Email,
			},
		}
	}
	utils.SendResponse(w, *response)
}
