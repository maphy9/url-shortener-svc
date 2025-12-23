-- +migrate Up
DROP TABLE IF EXISTS url_mapping;

CREATE TABLE url_mapping
(
    original_url TEXT NOT NULL UNIQUE,
    shortened_url TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS url_mapping_index ON url_mapping(original_url, shortened_url);
-- +migrate Down
