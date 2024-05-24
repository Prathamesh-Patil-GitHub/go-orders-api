package application

import (
	"context"
	"fmt"
	"net/http"
)

// Application type with a router of type http.Handler
type App struct {
	router http.Handler
}

// Constructor for application
func New() *App {
	app := &App{
		// This method loads all routes in the router which is actually a chi router (*chi.Mux)
		router: loadRoutes(),
	}

	return app
}

// This method will be called from the main file with the app object to start the server
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000", // Port of the http server
		Handler: a.router,
	}

	// This error is returned if the server doesn't start
	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
