package constants

const (
	// Apps
	AppDbMigrate    = "db-migrate"
	AppMendelServer = "mendel-server"

	// Environment/App class constants
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"

	// Databse constants
	SchemaMendelCore   = "mendel_core"
	DBInitQuery        = `SET search_path TO ` + SchemaMendelCore + `, public;`
	TablePlant         = "plant"
	TablePlantCultivar = "plant_cultivar"
	TablePlantSpecies  = "plant_species"
	TableUser          = "user"

	// Route params
	ID      = "id"
	ParamID = "/:id"

	// Routes
	RouteEnv           = "/env"
	RouteHealth        = "/health"
	RouteIndex         = "/"
	RoutePlant         = "/plant"
	RoutePlantCultivar = "/plant-cultivar"
	RoutePlantSpecies  = "/plant-species"
)
