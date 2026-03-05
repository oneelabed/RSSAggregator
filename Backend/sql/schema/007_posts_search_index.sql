-- +goose Up
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX IF NOT EXISTS idx_posts_search_trgm ON posts USING gin (title gin_trgm_ops, description gin_trgm_ops);

-- +goose Down
DROP INDEX IF EXISTS idx_posts_search_trgm;