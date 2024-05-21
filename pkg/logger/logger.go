package logger

import (
	"log"
	"log/slog"

	"github.com/jcmturner/gokrb5/v8/config"
)

type Logger struct {
	l *slog.Logger
}

func New(cfg *config.Config) *Logger {
	if cfg.Environment == "dev" {

	} else if cfg.Environment == "prod" {

	} else {
		log.Fatal("choose the environment")
	}
}
