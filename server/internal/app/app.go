// app.go

package app

import (
	"context"
	"database/sql"
	"fmt"

	// "log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/stdlib"

	"github.com/rs/zerolog/log"

	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/db"
	"github.com/kylep342/mendel/internal/handlers"
	"github.com/kylep342/mendel/internal/models/plants"
)

// App is the singleton struct with components to run mendel
type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

// Initialize creates the application's components.
func (a *App) Initialize(env *constants.EnvConfig) {
	var err error

	sqlURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		env.Database.Dialect,
		env.Database.User,
		env.Database.Password,
		env.Database.Host,
		env.Database.Port,
		env.Database.Name,
		env.Database.SSLMode,
	)

	// Postgres setup
	a.DB, err = sql.Open("pgx", sqlURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open database connection")
	}

	a.DB.SetMaxOpenConns(env.Database.MaxOpenConns)
	a.DB.SetMaxIdleConns(env.Database.MaxIdleConns)
	a.DB.SetConnMaxLifetime(env.Database.ConnMaxLifetime)

	// Ping the database to verify the connection.
	// Use a startup context for this operation.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = a.DB.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	_, err = a.DB.Exec(constants.DBInitQuery)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to set search_path")
	}

	log.Info().Msg("Database connection successful.")

	a.Router = gin.Default()
	a.InitializeRoutes(env)
}

// InitializeRoutes creates all endpoints for the api
func (a *App) InitializeRoutes(env *constants.EnvConfig) {
	// ... (your existing route initialization logic is perfect)
	// You can now pass the `env` to any handlers that might need it.
	a.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
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
	// plantSpeciesHandler.RegisterRoutes(a.Router, constants.RoutePlantSpecies)
	plantSpeciesHandler.RegisterRoutes(a.Router, constants.RoutePlantSpecies)

	plantCultivarHandler := handlers.NewCRUDHandler(
		a.DB,
		env,
		func() *plants.PlantCultivar { return &plants.PlantCultivar{} },
		func(d *sql.DB) db.CRUDTable[plants.PlantCultivar] {
			return &db.PlantCultivarTable{DB: d}
		},
	)
	// plantCultivarHandler.RegisterRoutes(a.Router, constants.RoutePlantCultivar)
	plantCultivarHandler.RegisterRoutes(a.Router, constants.RoutePlantCultivar)

	plantHandler := handlers.NewCRUDHandler(
		a.DB,
		env,
		func() *plants.Plant { return &plants.Plant{} },
		func(d *sql.DB) db.CRUDTable[plants.Plant] {
			return &db.PlantTable{DB: d}
		},
	)
	// plantHandler.RegisterRoutes(a.Router, constants.RoutePlant)
	plantHandler.RegisterRoutes(a.Router, constants.RoutePlant)
}

// Run starts the app and now includes graceful shutdown logic.
// It uses the timeout values from your environment configuration.
func (a *App) Run(env *constants.EnvConfig) {
	// RunServer(a.Router, env.Server.Host, env.Server.Port, env)
	RunServer(a.Router, env.Server.Host, env.Server.Port, env)
}
