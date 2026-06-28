package models

import (
	"time"

	"homepage/internal/database"
)

type CachedHTML struct {
	HTML       string
	Timestamp  time.Time
	LastAccess time.Time
}

type ContentProps struct {
	Content     []database.Content
	IsAdmin     bool
	IsEditing   bool
	ContentType string
	FullView    bool
}
