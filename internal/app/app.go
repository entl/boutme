package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.echo = echo.New()

	return a, nil
}

func (a *App) Run() error {

	fmt.Println("server is running...")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal("failed to start server", err)
	}

	return nil
}
