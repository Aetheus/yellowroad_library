ALTER TABLE chapter_paths
    ADD COLUMN effects JSONB NOT NULL DEFAULT '{}',
    ADD COLUMN requirements JSONB NOT NULL DEFAULT '{}';
