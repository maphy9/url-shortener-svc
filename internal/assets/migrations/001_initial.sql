-- +migrate Up
CREATE TABLE IF NOT EXISTS mapping
(
  url text NOT NULL UNIQUE,
  alias text NOT NULL UNIQUE,
  created_at timestamptz DEFAULT current_timestamp
);

CREATE SEQUENCE IF NOT EXISTS code_sequence;

-- +migrate Down
DROP SEQUENCE IF EXISTS code_sequence;

DROP TABLE IF EXISTS mapping CASCADE;