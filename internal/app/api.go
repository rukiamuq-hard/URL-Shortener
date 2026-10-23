package app

import (
	"URLS/internal/link"
	//"URLS/internal/monetization"
	"URLS/internal/platform/config"
	"URLS/internal/platform/postgres"
	"context"
	"time"

	"github.com/labstack/echo/v5"
)

type APIApp struct {
	echo *echo.Echo
	cfg  *config.Config
	pg   *postgres.Postgres
}

func NewApi(cfg *config.Config, pg *postgres.Postgres) *APIApp {
	e := echo.New()

	//adServ := monetization.NewMonetization()
	linkRepo := link.NewRepository(pg.DB)
	linkService := link.NewService(linkRepo)
	linkHandler := link.NewHandler(linkService)

	linkHandler.Register(e)
	return &APIApp{
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
