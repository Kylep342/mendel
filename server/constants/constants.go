package constants

const (
	SchemaMendelCore   = "mendel_core"
	DBInitQuery        = `SET search_path = ` + SchemaMendelCore + `, public;`
	TablePlantSpecies  = "plant_species"
	TablePlant         = "plant"
	TablePlantCultivar = "plant_cultivar"
	TableUser          = "user"

	QueryInsertPlantSpecies = `
		INSERT INTO ` + SchemaMendelCore + `.` + TablePlantSpecies + ` (name, taxon)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	QuerySelectPlantSpeciesByID = `
		SELECT id, name, taxon, created_at, updated_at
		FROM ` + SchemaMendelCore + `.` + TablePlantSpecies + `
		WHERE id = $1
	`
)
