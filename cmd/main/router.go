package main

import (
	"github.com/go-chi/chi"
	"plagChecker/internal/app"
)

func newRouter(a *app.App) chi.Router {
	r := chi.NewRouter()

	r.HandleFunc("/upload", a.UploadHandler)
	r.Get("/check/{name}", a.CheckStudentHandler)
	r.Get("/check/{labID}", a.CheckLabHandler)

	return r
}
