-- name: CreateAuthor :one
INSERT INTO authors (name, email) VALUES ($1, $2) RETURNING *;

-- name: ListAuthors :many
SELECT * FROM authors;

-- name: CreateBlogPost :one
INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING *;

-- name: ListBlogPosts :many
SELECT * FROM posts;

-- name: GetRecentBlogPosts :many
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
		LIMIT $1;

-- name: GetBlogPost :one
SELECT * FROM posts WHERE id = $1;

-- name: GetBlogPostByTitle :one
SELECT * FROM posts WHERE title = $1;

-- name: CreateComment :one
INSERT INTO comments (post_id, author_id, content) VALUES ($1, $2, $3) RETURNING *;

-- name: ListComments :many
SELECT * FROM comments WHERE post_id = $1;

-- name: CreateTag :one
INSERT INTO tags (name) VALUES ($1) RETURNING *;

-- name: AddTagToPost :exec
INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2);

-- name: GetPostTags :many
SELECT t.* FROM tags t
JOIN post_tags pt ON t.id = pt.tag_id
WHERE pt.post_id = $1;
