ALTER TABLE books
  ALTER COLUMN deleted_at TYPE timestamp without time zone,
  ALTER COLUMN created_at TYPE timestamp without time zone,
  ALTER COLUMN updated_at TYPE timestamp without time zone;

ALTER TABLE chapters
  ALTER COLUMN deleted_at TYPE timestamp without time zone,
  ALTER COLUMN created_at TYPE timestamp without time zone,
  ALTER COLUMN updated_at TYPE timestamp without time zone;

ALTER TABLE chapter_paths
    ALTER COLUMN deleted_at TYPE timestamp without time zone,
    ALTER COLUMN created_at TYPE timestamp without time zone,
    ALTER COLUMN updated_at TYPE timestamp without time zone;

ALTER TABLE users
  ALTER COLUMN deleted_at TYPE timestamp without time zone,
  ALTER COLUMN created_at TYPE timestamp without time zone,
  ALTER COLUMN updated_at TYPE timestamp without time zone;