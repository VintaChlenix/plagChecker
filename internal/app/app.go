package app

import (
	"go.uber.org/zap"
	"net/http"
	"plagChecker/internal/db"
)

type App struct {
	log *zap.SugaredLogger
	db  db.DB
}

func NewApp(log *zap.SugaredLogger, db db.DB) (*App, error) {
	return &App{
		log: log,
		db:  db,
	}, nil
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {

}
