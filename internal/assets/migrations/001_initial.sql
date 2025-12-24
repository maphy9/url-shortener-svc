-- +migrate Up
CREATE TABLE IF NOT EXISTS aliases
(
  url text NOT NULL UNIQUE,
  alias text NOT NULL UNIQUE,
  created_at timestamptz DEFAULT current_timestamp
);

CREATE SEQUENCE IF NOT EXISTS alias_sequence
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
CREATE OR REPLACE FUNCTION get_alias(v_url text)
RETURNS text
AS $body$
DECLARE
  res text := '';
BEGIN
  WITH set_alias AS (
    INSERT INTO aliases(url, alias)
    VALUES(v_url, nextval('alias_sequence'))
    ON CONFLICT (url) DO NOTHING
    RETURNING alias
  )
  SELECT alias FROM set_alias
  UNION ALL
  SELECT alias FROM aliases
  WHERE url = v_url LIMIT 1
  INTO res;

  RETURN res;
END;
$body$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION get_url(v_alias text)
RETURNS text
AS $body$
DECLARE
  res text := '';
BEGIN
  res := (SELECT url
  FROM aliases
  WHERE alias = v_alias);
  
  RETURN res;
END;
$body$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION IF EXISTS get_url(v_alias text);

DROP FUNCTION IF EXISTS get_alias(v_url text);

DROP FUNCTION IF EXISTS to_base62(val bigint);

DROP SEQUENCE IF EXISTS alias_sequence CASCADE;

DROP TABLE IF EXISTS url_aliases CASCADE;