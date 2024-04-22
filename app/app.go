package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	// Keeps the router decoupled from Chi
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}

	return nil

}
