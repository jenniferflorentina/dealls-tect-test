package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

var accessHandler = hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
	var event *zerolog.Event
	logBuilder := hlog.FromRequest(r)

	if status < 400 {
		event = logBuilder.Info()
	} else if status < 500 {
		event = logBuilder.Warn()
	} else {
		event = logBuilder.Error()
	}

	event.Str("method", r.Method).
		Int("statusCode", status).
		Dur("duration", duration).
		Msg(r.RequestURI)
})

func LogHandler(next http.Handler) http.Handler {
	return hlog.NewHandler(log.Logger)(accessHandler(next))
}

func SetupLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("::: %s", i)
	}

	log.Logger = log.Output(output)
}
