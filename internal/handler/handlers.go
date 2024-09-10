package handler

import (
	"homepage/internal/blog"
	"homepage/internal/database"
	"homepage/internal/markdown"
)

type Handler struct {
	DB          *database.Queries
	MD          markdown.MarkdownConverter
	BlogService *blog.BlogService
}

func NewHandler(db *database.Queries, md markdown.MarkdownConverter, blogService *blog.BlogService) *Handler {
	return &Handler{
		DB:          db,
		MD:          md,
		BlogService: blogService,
	}
}
