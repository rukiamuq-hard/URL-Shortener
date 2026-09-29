package main

import (
	"log"
)

func main() {
	a := app.New()
	err := a.Start()
	if err != nil {
		log.Fatal(err)
	}
}
