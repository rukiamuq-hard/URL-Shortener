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

func NewApp(ctx context.Context) (*App, error) {
	cfg := config.Load()
	pg := postgres.New()

	if err := pg.Connect(cfg.PostgresDSN, ctx); err != nil {
		return nil, err
	}

	apiApp := NewApi(cfg, pg)

	return &App{
		cfg:    cfg,
		apiApp: apiApp,
		pg:     pg,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	err := a.apiApp.Run(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Close() {
	a.pg.Close()
}
