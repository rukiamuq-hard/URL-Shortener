package app

import (
	"URLS/internal/platform/config"
	"URLS/internal/platform/postgres"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg    *config.Config
	apiApp *APIApp
	pg     *postgres.Postgress
}

func NewApp() *App {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg := config.Load()

	pg := postgres.New()
	if err := pg.Connect(cfg.PostgresDSN); err != nil {
		fmt.Println("Error PG:", err)
	}

	apiApp := NewApi(cfg)
	err := apiApp.Run(ctx)
	if err != nil {
		fmt.Println("Error", err)
	}
	return &App{
		cfg:    cfg,
		apiApp: apiApp,
		pg:     pg,
	}
}

func (a *App) Close() {
}
