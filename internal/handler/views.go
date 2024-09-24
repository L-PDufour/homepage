package handler

import (
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
