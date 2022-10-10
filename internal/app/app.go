package app

import (
	"net/http"
	"plagChecker/internal/db"
)

type App struct {
	db db.DB
}

func NewApp() (*App, error) {
	return &App{}, nil
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {

}
