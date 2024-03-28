package main

import (
	"github.com/entl/boutme/internal/app"
	"log"
)

func main() {
	a, err := app.New()

	if err != nil {
		log.Fatal("failed to start server", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}
