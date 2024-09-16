-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Step 1: Create the new enum type.
CREATE TYPE content_type AS ENUM ('blog', 'project', 'bio');

-- Step 2: Add the new column with the enum type.
ALTER TABLE content ADD COLUMN content_type content_type NOT NULL DEFAULT 'blog';

-- Step 3: If you already have a `type` column and want to migrate the data,
-- you can use this query (assuming `type` contains valid enum values):
UPDATE content SET content_type = CASE
  WHEN type = 'blog' THEN 'blog'::content_type
  WHEN type = 'project' THEN 'project'::content_type
  WHEN type = 'bio' THEN 'bio'::content_type
  ELSE 'blog'::content_type  -- Fallback for any other values
END;

-- Step 4: Drop the old `type` column.
ALTER TABLE content DROP COLUMN type;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

-- Step 1: Add the old `type` column back if needed (as a string).
ALTER TABLE content ADD COLUMN type TEXT;

-- Step 2: Migrate the data back from `content_type` to `type`.
UPDATE content SET type = content_type::TEXT;

-- Step 3: Drop the enum column.
ALTER TABLE content DROP COLUMN content_type;

-- Step 4: Drop the enum type.
DROP TYPE content_type;

