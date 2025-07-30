// app.go

package app

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rs/zerolog"

	"github.com/kylep342/mendel/internal/components/plant"
	"github.com/kylep342/mendel/internal/components/plant_cultivar"
	"github.com/kylep342/mendel/internal/components/plant_species"
	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/db"
	"github.com/kylep342/mendel/internal/handlers"
	"github.com/kylep342/mendel/pkg/responses"
)

// App is the singleton struct with components to run mendel
type App struct {
	DB     *pgxpool.Pool
	Logger zerolog.Logger
	Router *gin.Engine
}

// Run starts the app and performs a graceful shutdown.
func (a *App) Run(env *constants.EnvConfig) {
	RunServer(a.Router, env)
}

// Initialize creates the application's components.
func (a *App) Initialize(logger zerolog.Logger, env *constants.EnvConfig) {
	var err error

	a.Logger = logger

	ctx, cancel := context.WithTimeout(context.Background(), env.Server.ReadTimeout)
	defer cancel()

	// Postgres pgx setup
	pgConf, err := pgxpool.ParseConfig(env.DBUrl())
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to parse connection string")
	}
	a.DB, err = pgxpool.NewWithConfig(ctx, pgConf)
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("Failed to open database connection")
	}

	ctxPG, cancelPG := context.WithTimeout(context.Background(), env.Server.ReadTimeout)
	defer cancelPG()

	_, err = a.DB.Exec(ctxPG, constants.DBInitQuery)
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to set search_path")
	}

	// Router setup
	a.Router = gin.Default()
	a.setupMiddleware(env)
	a.InitializeRoutes(env)
}

func (a *App) setupMiddleware(env *constants.EnvConfig) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{env.App.WebHost}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	a.Router.Use(cors.New(config))
}

// InitializeRoutes creates all endpoints for the api
func (a *App) InitializeRoutes(env *constants.EnvConfig) {
	a.Logger.Info().Msg("Initializing routes")

	a.Router.GET("/", func(c *gin.Context) {
		responses.RespondData(c, "ok", http.StatusOK)
	})

	internalHandler := handlers.NewInternalHandler(a.DB, env)
	internalHandler.RegisterRoutes(a.Router, constants.RouteIndex)

	plantSpeciesHandler := handlers.NewCRUDHandler(
		a.DB,
		env,
		func() *plant_species.PlantSpecies { return &plant_species.PlantSpecies{} },
		func(p *pgxpool.Pool) db.CRUDTable[plant_species.PlantSpecies] {
			return &plant_species.Store{conn: p}
		},
	)
	plantSpeciesHandler.RegisterRoutes(a.Router, constants.RoutePlantSpecies)

	plantCultivarHandler := handlers.NewCRUDHandler(
		a.DB,
		env,
		func() *plant_cultivar.PlantCultivar { return &plant_cultivar.PlantCultivar{} },
		func(p *pgxpool.Pool) db.CRUDTable[plant_cultivar.PlantCultivar] {
			return &plant_cultivar.Store{conn: p}
		},
	)
	plantCultivarHandler.RegisterRoutes(a.Router, constants.RoutePlantCultivar)

	plantHandler := handlers.NewCRUDHandler(
		a.DB,
		env,
		func() *plant.Plant { return &plant.Plant{} },
		func(p *pgxpool.Pool) db.CRUDTable[plant.Plant] {
			return &plant.Store{conn: p}
		},
	)
	plantHandler.RegisterRoutes(a.Router, constants.RoutePlant)
	a.Logger.Info().Msg("Routes initialized")
}
