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
	CurrentPath string
}

type ContentItemProps struct {
	Content   database.Content
	IsAdmin   bool
	IsEditing bool
}

type ContentProps struct {
	Content     database.Content
	IsAdmin     bool
	IsEditing   bool
	ContentType database.ContentType
}

type ListProps struct {
	Contents    []database.Content
	IsAdmin     bool
	ContentType database.ContentType
	CurrentPath string
}
