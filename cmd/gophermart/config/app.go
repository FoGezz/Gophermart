package config

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type App struct {
	DB     *sqlx.DB
	Logger *zap.Logger
}

func NewApp(config *Config, db *sqlx.DB) *App {
	var logger *zap.Logger
	if config.DevMode {
		logger = zap.Must(zap.NewDevelopment())
		zap.ReplaceGlobals(logger)
	} else {
		logger = zap.Must(zap.NewProduction())
		zap.ReplaceGlobals(logger)
	}
	return &App{
		Logger: logger,
		DB:     db,
	}
}
