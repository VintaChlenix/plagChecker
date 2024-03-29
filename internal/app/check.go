package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"plagChecker/internal/dto"
)

func (a *App) CheckStudentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := chi.URLParam(r, "name")

	response, err := a.checkStudentHandler(ctx, name)
	if err != nil {
		a.log.Errorf("failed to check student: %w", err)
		w.Write([]byte("error, check logs"))
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/check.html")
	if err != nil {
		a.log.Errorf("failed to parse template file: %w", err)
		return
	}
	tmpl.Execute(w, response.StudentCheckResults)
}

func (a *App) checkStudentHandler(ctx context.Context, name string) (*dto.CheckStudentResponse, error) {
	studentLabs, err := a.db.SelectStudentLabs(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to select student labs: %w", err)
	}
	return &dto.CheckStudentResponse{StudentCheckResults: studentLabs}, nil
}

func (a *App) CheckLabHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	labID := chi.URLParam(r, "labID")

	response, err := a.checkLabHandler(ctx, labID)
	if err != nil {
		a.log.Errorf("failed to check lab: %w", err)
		w.Write([]byte("error, check logs"))
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/checkLab.html")
	if err != nil {
		a.log.Errorf("failed to parse template file: %w", err)
		return
	}
	tmpl.Execute(w, response.LabCheckResults)
}

func (a *App) checkLabHandler(ctx context.Context, labID string) (*dto.CheckLabResponse, error) {
	labSendings, err := a.db.SelectLabSendings(ctx, labID)
	if err != nil {
		return nil, fmt.Errorf("failed to select lab sendings: %w", err)
	}
	return &dto.CheckLabResponse{LabCheckResults: labSendings}, nil
}
