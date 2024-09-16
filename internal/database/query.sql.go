// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const createContent = `-- name: CreateContent :one
INSERT INTO content (content_type, title, markdown, image_url, link, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
RETURNING id, content_type, title, markdown, image_url, link, created_at, updated_at
`

type CreateContentParams struct {
	ContentType ContentType
	Title       string
	Markdown    sql.NullString
	ImageUrl    sql.NullString
	Link        sql.NullString
}

type CreateContentRow struct {
	ID          int32
	ContentType ContentType
	Title       string
	Markdown    sql.NullString
	ImageUrl    sql.NullString
	Link        sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

func (q *Queries) CreateContent(ctx context.Context, arg CreateContentParams) (CreateContentRow, error) {
	row := q.db.QueryRowContext(ctx, createContent,
		arg.ContentType,
		arg.Title,
		arg.Markdown,
		arg.ImageUrl,
		arg.Link,
	)
	var i CreateContentRow
	err := row.Scan(
		&i.ID,
		&i.ContentType,
		&i.Title,
		&i.Markdown,
		&i.ImageUrl,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteContent = `-- name: DeleteContent :exec
DELETE FROM content WHERE id = $1
`

func (q *Queries) DeleteContent(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteContent, id)
	return err
}

const getContentById = `-- name: GetContentById :one
SELECT id, title, markdown, image_url, link, created_at, updated_at, content_type FROM content WHERE id = $1
`

// TODO
func (q *Queries) GetContentById(ctx context.Context, id int32) (Content, error) {
	row := q.db.QueryRowContext(ctx, getContentById, id)
	var i Content
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Markdown,
		&i.ImageUrl,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ContentType,
	)
	return i, err
}

const getContentByTitle = `-- name: GetContentByTitle :one
SELECT id, title, markdown, image_url, link, created_at, updated_at, content_type FROM content WHERE content_type = $1 AND title = $2
`

type GetContentByTitleParams struct {
	ContentType ContentType
	Title       string
}

func (q *Queries) GetContentByTitle(ctx context.Context, arg GetContentByTitleParams) (Content, error) {
	row := q.db.QueryRowContext(ctx, getContentByTitle, arg.ContentType, arg.Title)
	var i Content
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Markdown,
		&i.ImageUrl,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ContentType,
	)
	return i, err
}

const getContentsByType = `-- name: GetContentsByType :many
SELECT id, title, markdown, image_url, link, created_at, updated_at, content_type FROM content WHERE content_type = $1
`

func (q *Queries) GetContentsByType(ctx context.Context, contentType ContentType) ([]Content, error) {
	rows, err := q.db.QueryContext(ctx, getContentsByType, contentType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Content
	for rows.Next() {
		var i Content
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Markdown,
			&i.ImageUrl,
			&i.Link,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ContentType,
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

const updateContent = `-- name: UpdateContent :one
UPDATE content
SET content_type = $1, title = $2, markdown = $3, image_url = $4, link = $5, updated_at = NOW()
WHERE id = $6
RETURNING id, title, markdown, image_url, link, created_at, updated_at, content_type
`

type UpdateContentParams struct {
	ContentType ContentType
	Title       string
	Markdown    sql.NullString
	ImageUrl    sql.NullString
	Link        sql.NullString
	ID          int32
}

func (q *Queries) UpdateContent(ctx context.Context, arg UpdateContentParams) (Content, error) {
	row := q.db.QueryRowContext(ctx, updateContent,
		arg.ContentType,
		arg.Title,
		arg.Markdown,
		arg.ImageUrl,
		arg.Link,
		arg.ID,
	)
	var i Content
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Markdown,
		&i.ImageUrl,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ContentType,
	)
	return i, err
}
