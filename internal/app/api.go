package app

import (
	"URLS/internal/link"
	"URLS/internal/monetization"
	"URLS/internal/platform/config"
	"context"
	"github.com/labstack/echo/v5"
	"time"
)

type APIApp struct {
	echo *echo.Echo
	cfg  *config.Config
}

func NewApi(cfg *config.Config) *APIApp {
	e := echo.New()

	adServ := monetization.NewMonetization()

	linkService := link.NewService(adServ)
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
