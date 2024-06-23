package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

// add the owner of this method(receiver) to the function signature
// similar to .this in JavaScript
func(a *App) Start(ctx context.Context) error {
	// first define the http server
	server := &http.Server{
		Addr: ":8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server failed to start: %w", err)
}
	return nil
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

