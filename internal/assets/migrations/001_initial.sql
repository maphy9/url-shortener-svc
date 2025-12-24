-- +migrate Up
CREATE TABLE IF NOT EXISTS url_mapping
(
    url TEXT NOT NULL UNIQUE,
    code TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT current_timestamp
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

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION from_base62(val text)
RETURNS bigint
AS $body$
DECLARE
  chars text := '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
  res bigint := 0;
  i int;
BEGIN
  FOR i IN 1..length(val) LOOP
    res := res * 62 + (strpos(chars, substr(val, i, 1)) - 1);
  END LOOP;
  RETURN res;
END;
$body$ LANGUAGE plpgsql IMMUTABLE;
-- +migrate StatementEnd

-- +migrate Down
