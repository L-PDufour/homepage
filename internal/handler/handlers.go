package handler

import (
	"homepage/internal/blog"
	"homepage/internal/database"
)

type Handler struct {
	DB          *database.Queries
	BlogService *blog.BlogService
}

func NewHandler(db *database.Queries, blogService *blog.BlogService) *Handler {
	return &Handler{
		DB:          db,
		BlogService: blogService,
	}
}
