DROP SCHEMA IF EXISTS mendel_core;

CREATE SCHEMA IF NOT EXISTS mendel_core;

CREATE OR REPLACE FUNCTION trigger_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE mendel_core.users (
    id UUID DEFAULT gen_random_uuid(),
    username VARCHAR(49) NOT NULL,
    enabled boolean NOT NULL DEFAULT FALSE,
    web_settings JSONB NOT NULL,
    last_login timestamp,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

BEGIN;
DROP TRIGGER IF EXISTS users_update_timestamp ON users;

CREATE TRIGGER users_update_timestamp
BEFORE UPDATE ON mendel_core.users
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();
COMMIT;
