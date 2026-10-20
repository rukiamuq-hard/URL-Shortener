package main

import (
	"URLS/internal/app"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	a := app.NewApp()
	err := a.Run(ctx)
	defer a.Close()
	if err != nil {
		log.Fatal(err)
	}
}
