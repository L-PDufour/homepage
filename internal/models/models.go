package models

import (
	"homepage/internal/database"
	"time"
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
	ContentType database.ContentType
	FullView    bool
}
