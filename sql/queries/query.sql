-- name: CreateContent :one
INSERT INTO content (content_type, title, markdown, image_url, link, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
RETURNING id, content_type, title, markdown, image_url, link, created_at, updated_at;

-- name: GetContentById :one
--TODO
SELECT * FROM content WHERE id = $1;

-- name: GetContentByTitle :one
SELECT * FROM content WHERE content_type = $1 AND title = $2;

-- name: GetContentsByType :many
SELECT * FROM content WHERE content_type = $1 ORDER BY updated_at DESC;

-- name: UpdateContent :one
UPDATE content
SET content_type = $1, title = $2, markdown = $3, image_url = $4, link = $5, updated_at = NOW()
WHERE id = $6
RETURNING *;

-- name: DeleteContent :exec
DELETE FROM content WHERE id = $1;

