package handler

import (
	"net/http"

	"homepage/internal/database"
	"homepage/internal/models"
	"homepage/internal/views"
)

func (h *Handler) AdminAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.AdminAuthPage().Render(w)
	}
}

// func (h *Handler) ServeResume() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		views.Re().Render(w)
// 	}
// }

func (h *Handler) renderContentList(w http.ResponseWriter, r *http.Request, props models.ContentProps) {
	views.Base(views.ContentList(props)).Render(w)
}

func (h *Handler) renderNewForm(w http.ResponseWriter, r *http.Request) {
	views.Base(views.NewForm()).Render(w)
}

func (h *Handler) renderUpdateForm(w http.ResponseWriter, r *http.Request, props database.Content) {
	views.Base(views.EditForm(props)).Render(w)
}
