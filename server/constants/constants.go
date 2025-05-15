package constants

const (
	SchemaMendelCore   = "mendel_core"
	DBInitQuery        = `SET search_path = ` + SchemaMendelCore + `, public;`
	TablePlantSpecies  = "plant_species"
	TablePlant         = "plant"
	TablePlantCultivar = "plant_cultivar"
	TableUser          = "user"

	RoutePlant         = "/plant"
	RoutePlantCultivar = "/plant-cultivar"
	RoutePlantSpecies  = "/plant-species"
)
