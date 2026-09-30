package app

import (
	"URLS/internal/platform/config"
	"URLS/internal/platform/postgres"
	"fmt"
	//"github.com/labstack/echo/v5"
)

type APIApp struct {
	pdb *postgres.Postgress
	//echo *echo.Echo
	cfg *config.Config
}

type Handlers struct {
}

func NewApi(cfg *config.Config) *APIApp {

	return &APIApp{
		pdb: postgres.New(),
		cfg: cfg,
	}
}

func (aApp *APIApp) Close() {

	fmt.Println("closed\n", aApp.cfg)
}
