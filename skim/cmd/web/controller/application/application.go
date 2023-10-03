package application

import (
	"log/slog"
	"os"
)

type Application struct {
	Logger *slog.Logger
}

func NewApplication() (*Application, error) {
	payload := new(Application)
	handler := slog.NewJSONHandler(os.Stdout, nil)
	payload.Logger = slog.New(handler)

	return payload, nil
}
