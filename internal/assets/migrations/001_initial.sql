-- +migrate Up
CREATE TABLE IF NOT EXISTS mapping
(
  url text NOT NULL UNIQUE,
  alias text NOT NULL UNIQUE,
  created_at timestamptz DEFAULT current_timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS mapping CASCADE;