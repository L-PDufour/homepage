-- +goose Up
-- +goose StatementBegin
INSERT INTO authors (id, name, email)
VALUES
  (1, 'Author One', 'author1@example.com'),
  (2, 'Author Two', 'author2@example.com'),
  (3, 'Author Three', 'author3@example.com'),
  (4, 'Author Four', 'author4@example.com');

INSERT INTO posts (title, content, author_id)
VALUES
  ('First Post', 'This is the content of the first post.', 1),
  ('Second Post', 'Here is some content for the second post.', 2),
  ('Third Post', 'Content for the third post goes here.', 1),
  ('Fourth Post', 'Another post content for testing.', 4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM authors
WHERE name IN(
'Author One',
'Author Two',
'Author Three',
'Author Fourth',
);
DELETE FROM posts
WHERE title IN (
  'First Post',
  'Second Post',
  'Third Post',
  'Fourth Post'
);
-- +goose StatementEnd

