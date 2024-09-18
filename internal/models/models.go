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

type ContentViewProps struct {
	Contents    []database.Content
	ContentType database.ContentType
	IsAdmin     bool
	IsEditing   bool
}

type ContentItemProps struct {
	Content   database.Content
	IsAdmin   bool
	IsEditing bool
}
