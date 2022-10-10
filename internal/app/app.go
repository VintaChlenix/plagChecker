package app

import (
	"net/http"
	"plagChecker/internal/db"
)

type App struct {
	db db.DB
}

func NewApp(db db.DB) (*App, error) {
	return &App{db: db}, nil
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {

}
