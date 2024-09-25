package handler

import (
	"homepage/internal/database"
	"homepage/internal/models"
	"homepage/internal/views"
	"net/http"
)

func (h *Handler) AdminAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.AdminAuthPage().Render(r.Context(), w)
	}
}

func (h *Handler) ServeResume() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		views.ResumeTemplate().Render(r.Context(), w)
	}
}

func (h *Handler) renderContentList(w http.ResponseWriter, r *http.Request, props models.ContentProps) {
	views.Base(views.ContentList(props)).Render(r.Context(), w)
}

func (h *Handler) renderNewForm(w http.ResponseWriter, r *http.Request) {
	views.Base(views.NewForm()).Render(r.Context(), w)
}

func (h *Handler) renderUpdateForm(w http.ResponseWriter, r *http.Request, props database.Content) {
	views.Base(views.EditForm(props)).Render(r.Context(), w)
}
