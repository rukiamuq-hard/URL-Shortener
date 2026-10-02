package app

import (
	"URLS/internal/link"
	"URLS/internal/platform/config"
	"URLS/internal/platform/postgres"
	"context"
	"github.com/labstack/echo/v5"
	"time"
)

type APIApp struct {
	pdb  *postgres.Postgress
	echo *echo.Echo
	cfg  *config.Config
}

type Handlers struct {
}

func NewApi(cfg *config.Config) *APIApp {
	e := echo.New()
	linkHandler := link.NewHandler()

	linkHandler.Register(e)
	return &APIApp{
		pdb:  postgres.New(),
		echo: e,
		cfg:  cfg,
	}
}

func (a *APIApp) Run(ctx context.Context) error {
	sc := echo.StartConfig{
		Address:         a.cfg.HTTPAddr,
		GracefulTimeout: 5 * time.Second,
	}
	return sc.Start(ctx, a.echo)
}
