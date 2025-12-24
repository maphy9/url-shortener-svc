-- +migrate Up
CREATE TABLE IF NOT EXISTS url_mapping
(
  url text NOT NULL UNIQUE,
  code text NOT NULL UNIQUE,
  created_at timestamptz DEFAULT current_timestamp
);

CREATE SEQUENCE IF NOT EXISTS code_sequence
INCREMENT BY 1
START WITH 1;

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION to_base62(val bigint)
RETURNS text
AS $body$
DECLARE
  chars text := '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
  res text := '';
BEGIN
  IF val = 0 THEN RETURN '0'; END IF;
  WHILE val > 0 LOOP
    res := substr(chars, (val % 62)::int + 1, 1) || res;
    val := val / 62;
  END LOOP;
  RETURN res;
END;
$body$ LANGUAGE plpgsql IMMUTABLE;
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION IF EXISTS to_base62(val bigint);

DROP SEQUENCE IF EXISTS code_sequence CASCADE;

DROP TABLE IF EXISTS url_mapping CASCADE;