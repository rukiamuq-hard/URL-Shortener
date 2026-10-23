package app

import (
	"URLS/internal/platform/config"
	"URLS/internal/platform/postgres"
	"context"
)

type App struct {
	cfg    *config.Config
	apiApp *APIApp
	pg     *postgres.Postgres
}

func NewApp() *App {
	cfg := config.Load()
	pg := postgres.New()
	apiApp := NewApi(cfg, pg)

	return &App{
		cfg:    cfg,
		apiApp: apiApp,
		pg:     pg,
	}
}

func (a *App) Run(ctx context.Context) error {
	if err := a.pg.Connect(a.cfg.PostgresDSN, ctx); err != nil {
		return err
	}

	err := a.apiApp.Run(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Close() {
	a.pg.Close()
}
