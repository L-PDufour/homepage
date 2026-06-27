CREATE TABLE IF NOT EXISTS content (
    id          INTEGER PRIMARY KEY,
    content_type TEXT NOT NULL CHECK (content_type IN ('blog', 'project', 'bio')),
    title       TEXT NOT NULL,
    content     TEXT,
    image_url   TEXT,
    link        TEXT,
    created_at  TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    updated_at  TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    UNIQUE (content_type, title)
);
