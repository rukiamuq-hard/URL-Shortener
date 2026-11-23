package app

import (
	"URLS/internal/platform/config"
	"URLS/internal/platform/postgres"
	"context"
	"log/slog"
)

type App struct {
	cfg    *config.Config
	apiApp *APIApp
	pg     *postgres.Postgres
}

func NewApp(ctx context.Context) (*App, error) {
	cfg := config.Load()
	pg := postgres.New()

	if err := pg.Connect(ctx, cfg.PostgresDSN); err != nil {
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
	err := a.pg.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}
