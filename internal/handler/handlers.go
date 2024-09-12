package handler

import (
	"homepage/internal/database"
)

type Handler struct {
	DB *database.Queries
}

func NewHandler(db *database.Queries) *Handler {
	return &Handler{
		DB: db,
	}
}
