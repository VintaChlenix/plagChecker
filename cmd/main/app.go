package main

import (
	"fmt"
	"net/http"
	"plagChecker/internal/app"
)

func run() error {
	/*dbconn, err := postgres.NewClient("") //todo: make config file
	if err != nil {
		return fmt.Errorf("failed to make db client: %w", err)
	}*/
	checkerApp, err := app.NewApp()
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}
	if err := http.ListenAndServe(":8080", newRouter(checkerApp)); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

func main() {
	//todo: need logs
	if err := run(); err != nil {
		panic(err.Error())
	}
}
