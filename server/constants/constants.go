package constants

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"

	DBInitQuery        = `SET search_path = ` + SchemaMendelCore + `, public;`
	SchemaMendelCore   = "mendel_core"
	TablePlant         = "plant"
	TablePlantCultivar = "plant_cultivar"
	TablePlantSpecies  = "plant_species"
	TableUser          = "user"

	RouteHealth        = "/health"
	RouteIndex         = "/"
	RoutePlant         = "/plant"
	RoutePlantCultivar = "/plant-cultivar"
	RoutePlantSpecies  = "/plant-species"
)
