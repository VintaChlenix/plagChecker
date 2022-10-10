package main

import (
	"github.com/go-chi/chi"
	"plagChecker/internal/app"
)

func newRouter(a *app.App) chi.Router {
	r := chi.NewRouter()

	r.Get("/", a.IndexHandler)
	r.HandleFunc("/upload", a.UploadHandler)

	return r
}
