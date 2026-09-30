package main

import (
	"URLS/internal/app"
)

func main() {
	a := app.NewApp()
	defer a.Close()
}
