-- name: CreateContent :one
INSERT INTO content (content_type, title, content, image_url, link)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetContentById :one
SELECT * FROM content WHERE id = ?;

-- name: GetContentByTitle :one
SELECT * FROM content WHERE content_type = ? AND title = ?;

-- name: GetContentsByType :many
SELECT * FROM content WHERE content_type = ? ORDER BY updated_at DESC;

-- name: UpdateContent :one
UPDATE content
SET content_type = ?, title = ?, content = ?, image_url = ?, link = ?, updated_at = strftime('%Y-%m-%dT%H:%M:%SZ', 'now')
WHERE id = ?
RETURNING *;

-- name: DeleteContent :exec
DELETE FROM content WHERE id = ?;
