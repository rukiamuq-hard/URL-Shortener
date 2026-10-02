package app

import (
	"URLS/internal/platform/config"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg    *config.Config
	apiApp *APIApp
}

func NewApp() *App {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	cfg := config.Load()
	apiApp := NewApi(cfg)

	err := apiApp.Run(ctx)
	if err != nil {
		fmt.Println("Error", err)
	}
	return &App{
		cfg:    cfg,
		apiApp: apiApp,
	}
}

func (a *App) Close() {
}
