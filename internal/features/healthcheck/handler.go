package healthcheck

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

type handler struct{}

type Response struct {
	BuildTime string `json:"build_time"`
}

func New() httprouter.Handle {
	return handler{}.Handle
}

func (h handler) Handle(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	buildTime := viper.GetString("buildTime")
	response := Response{
		BuildTime: buildTime,
	}

	res, _ := json.Marshal(response)
	_, err := w.Write([]byte(res))

	if err != nil {
		return
	}
}
