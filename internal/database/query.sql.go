// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const addTagToPost = `-- name: AddTagToPost :exec
INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)
`

type AddTagToPostParams struct {
	PostID int32
	TagID  int32
}

func (q *Queries) AddTagToPost(ctx context.Context, arg AddTagToPostParams) error {
	_, err := q.db.ExecContext(ctx, addTagToPost, arg.PostID, arg.TagID)
	return err
}

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at, updated_at
`

type CreateAuthorParams struct {
	Name  string
	Email string
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Email)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createBlogPost = `-- name: CreateBlogPost :one
INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING id, title, content, author_id, created_at, updated_at
`

type CreateBlogPostParams struct {
	Title    string
	Content  string
	AuthorID sql.NullInt32
}

func (q *Queries) CreateBlogPost(ctx context.Context, arg CreateBlogPostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createBlogPost, arg.Title, arg.Content, arg.AuthorID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createComment = `-- name: CreateComment :one
INSERT INTO comments (post_id, author_id, content) VALUES ($1, $2, $3) RETURNING id, post_id, author_id, content, created_at, updated_at
`

type CreateCommentParams struct {
	PostID   sql.NullInt32
	AuthorID sql.NullInt32
	Content  string
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.PostID, arg.AuthorID, arg.Content)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.AuthorID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTag = `-- name: CreateTag :one
INSERT INTO tags (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) CreateTag(ctx context.Context, name string) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createTag, name)
	var i Tag
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getBlogPost = `-- name: GetBlogPost :one
SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE id = $1
`

func (q *Queries) GetBlogPost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getBlogPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBlogPostByTitle = `-- name: GetBlogPostByTitle :one
SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE title = $1
`

func (q *Queries) GetBlogPostByTitle(ctx context.Context, title string) (Post, error) {
	row := q.db.QueryRowContext(ctx, getBlogPostByTitle, title)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostTags = `-- name: GetPostTags :many
SELECT t.id, t.name FROM tags t
JOIN post_tags pt ON t.id = pt.tag_id
WHERE pt.post_id = $1
`

func (q *Queries) GetPostTags(ctx context.Context, postID int32) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, getPostTags, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecentBlogPosts = `-- name: GetRecentBlogPosts :many
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
		LIMIT $1
`

func (q *Queries) GetRecentBlogPosts(ctx context.Context, limit int32) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getRecentBlogPosts, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, email, created_at, updated_at FROM authors
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBlogPosts = `-- name: ListBlogPosts :many
SELECT id, title, content, author_id, created_at, updated_at FROM posts
`

func (q *Queries) ListBlogPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listBlogPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listComments = `-- name: ListComments :many
SELECT id, post_id, author_id, content, created_at, updated_at FROM comments WHERE post_id = $1
`

func (q *Queries) ListComments(ctx context.Context, postID sql.NullInt32) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listComments, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.AuthorID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
