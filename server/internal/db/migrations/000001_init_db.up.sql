DROP SCHEMA IF EXISTS mendel_core CASCADE;

CREATE SCHEMA IF NOT EXISTS mendel_core;

CREATE OR REPLACE FUNCTION mendel_core.trigger_update_timestamp()
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
DROP TRIGGER IF EXISTS users_update_timestamp ON mendel_core.users;

CREATE TRIGGER users_update_timestamp
BEFORE UPDATE ON mendel_core.users
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();
COMMIT;

CREATE TABLE
    IF NOT EXISTS mendel_core.plant_species (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name TEXT NOT NULL,
        taxon TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT now (),
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT now ()
    );

BEGIN;

DROP TRIGGER IF EXISTS plant_species_update_timestamp ON mendel_core.plant_species;

CREATE TRIGGER plant_species_update_timestamp BEFORE
UPDATE ON mendel_core.plant_species FOR EACH ROW EXECUTE PROCEDURE trigger_update_timestamp ();

COMMIT;

CREATE TABLE
    IF NOT EXISTS mendel_core.plant_cultivar (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        species_id UUID NOT NULL REFERENCES mendel_core.plant_species (id) ON DELETE CASCADE ON UPDATE RESTRICT,
        name TEXT NOT NULL,
        cultivar TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT now (),
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT now (),
        genetics JSONB NOT NULL DEFAULT '{}'
    );

BEGIN;

DROP TRIGGER IF EXISTS plant_cultivar_update_timestamp ON mendel_core.plant_cultivar;

CREATE TRIGGER plant_cultivar_update_timestamp BEFORE
UPDATE ON mendel_core.plant_cultivar FOR EACH ROW EXECUTE PROCEDURE trigger_update_timestamp ();

COMMIT;


CREATE TABLE
    IF NOT EXISTS mendel_core.plant (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        cultivar_id UUID NOT NULL REFERENCES mendel_core.plant_cultivar (id) ON DELETE CASCADE ON UPDATE RESTRICT,
        species_id UUID NOT NULL REFERENCES mendel_core.plant_species (id) ON DELETE CASCADE ON UPDATE RESTRICT,
        seed_id UUID REFERENCES mendel_core.plant (id),
        pollen_id UUID REFERENCES mendel_core.plant (id),
        generation INT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT now (),
        updated_at TIMESTAMP WITH TIME ZONE,
        genetics JSONB NOT NULL DEFAULT '{}',
        labels JSONB NOT NULL DEFAULT '{}'
    );

BEGIN;

DROP TRIGGER IF EXISTS plant_cultivar_update_timestamp ON mendel_core.plant_cultivar;

CREATE TRIGGER plant_cultivar_update_timestamp BEFORE
UPDATE ON mendel_core.plant_cultivar FOR EACH ROW EXECUTE PROCEDURE trigger_update_timestamp ();

COMMIT;
