package application

import (
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	Logger *slog.Logger
}

func NewApplication() (*Application, error) {
	payload := new(Application)
	options := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	payload.Logger = slog.New(handler)

	return payload, nil
}

func (app *Application) catchHandlerErr(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
	app.Logger.Error("controller error", slog.With(err))
}
