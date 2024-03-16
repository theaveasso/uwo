package main

import (
	"fmt"
	"net/http"
)

type app struct{}

func newApp() *app {
    return &app{}
}

func (app *app) start() error {
    server := &http.Server{
        Addr: ":3333",
        Handler: setupRoutes(),
    }

    err := server.ListenAndServe() 
    if err != nil {
        return fmt.Errorf("failed to start server: %w", err)
    }

    return nil
}
