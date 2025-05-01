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
