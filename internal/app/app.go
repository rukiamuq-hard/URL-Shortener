package app

import (
	"URLS/internal/platform/config"
)

type App struct {
	cfg    *config.Config
	apiApp *APIApp
}

func NewApp() *App {
	cfg := config.Load()
	apiApp := NewApi(cfg)

	return &App{
		cfg:    cfg,
		apiApp: apiApp,
	}
}

func (a *App) Close() {
	defer a.apiApp.Close()
}
