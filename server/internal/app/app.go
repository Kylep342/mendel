// app.go

package app

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rs/zerolog"

	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/db"
	"github.com/kylep342/mendel/internal/handlers"
	"github.com/kylep342/mendel/internal/models/plants"
	"github.com/kylep342/mendel/pkg/responses"
)

// App is the singleton struct with components to run mendel
type App struct {
	DB     *sql.DB
	DBPG   *pgxpool.Pool
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

	// Postgres setup
	a.DB, err = sql.Open("pgx", env.DBUrl())
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("Failed to open database connection")
	}

	a.DB.SetMaxOpenConns(env.Database.MaxOpenConns)
	a.DB.SetMaxIdleConns(env.Database.MaxIdleConns)
	a.DB.SetConnMaxLifetime(env.Database.ConnMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), env.Server.ReadTimeout)
	defer cancel()

	if err = a.DB.PingContext(ctx); err != nil {
		a.Logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	_, err = a.DB.Exec(constants.DBInitQuery)
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to set search_path")
	}

	a.Logger.Info().Msg("Database connection successful")

	// Postgres pgx setup
	a.DBPG, err = pgxpool.Connect(ctx, env.DBUrl())
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("Failed to open database connection")
	}

	ctxPG, cancelPG := context.WithTimeout(context.Background(), env.Server.ReadTimeout)
	defer cancelPG()

	_, err = a.DBPG.Exec(ctxPG, constants.DBInitQuery)
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
		func() *plants.PlantSpecies { return &plants.PlantSpecies{} },
		func(d *sql.DB) db.CRUDTable[plants.PlantSpecies] {
			return &db.PlantSpeciesTable{DB: d}
		},
	)
	plantSpeciesHandler.RegisterRoutes(a.Router, constants.RoutePlantSpecies)

	plantCultivarHandler := handlers.NewCRUDHandler(
		a.DB,
		env,
		func() *plants.PlantCultivar { return &plants.PlantCultivar{} },
		func(d *sql.DB) db.CRUDTable[plants.PlantCultivar] {
			return &db.PlantCultivarTable{DB: d}
		},
	)
	plantCultivarHandler.RegisterRoutes(a.Router, constants.RoutePlantCultivar)

	plantHandler := handlers.NewCRUDHandler(
		a.DB,
		env,
		func() *plants.Plant { return &plants.Plant{} },
		func(d *sql.DB) db.CRUDTable[plants.Plant] {
			return &db.PlantTable{DB: d}
		},
	)
	plantHandler.RegisterRoutes(a.Router, constants.RoutePlant)
	a.Logger.Info().Msg("Routes initialized")
}
