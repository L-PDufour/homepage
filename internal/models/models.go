package models

import (
	"time"
)

// Author represents a blog post author
type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Post represents a blog post
type Post struct {
	ID        int32     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	Author    *Author   `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Comment represents a comment on a blog post
type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	AuthorID  int       `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tag represents a tag that can be applied to blog posts
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PostTag represents the many-to-many relationship between posts and tags
type PostTag struct {
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}
