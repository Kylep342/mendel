// app.go

package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/stdlib"

	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/db"
	"github.com/kylep342/mendel/internal/handlers"
	"github.com/kylep342/mendel/internal/models/plants"
	"github.com/kylep342/mendel/pkg/responses"
)

// App is the singleton struct with components to run mendel
type App struct {
	DB     *sql.DB
	Router chi.Router
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
		log.Fatalf("FATAL: Failed to open database connection: %v", err)
	}

	a.DB.SetMaxOpenConns(env.Database.MaxOpenConns)
	a.DB.SetMaxIdleConns(env.Database.MaxIdleConns)
	a.DB.SetConnMaxLifetime(env.Database.ConnMaxLifetime)

	// Ping the database to verify the connection.
	// Use a startup context for this operation.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = a.DB.PingContext(ctx); err != nil {
		log.Fatalf("FATAL: Failed to connect to database: %v", err)
	}

	_, err = a.DB.Exec(constants.DBInitQuery)
	if err != nil {
		log.Fatal("FATAL: failed to set search_path: %w", err)
	}

	log.Println("INFO: Database connection successful.")

	a.Router = chi.NewRouter()
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)

	a.InitializeRoutes(env) // Pass config to routes if needed
}

// InitializeRoutes creates all endpoints for the api
func (a *App) InitializeRoutes(env *constants.EnvConfig) {
	// ... (your existing route initialization logic is perfect)
	// You can now pass the `env` to any handlers that might need it.
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responses.RespondWithData(w, http.StatusOK, "ok")
	})

	internalHandler := handlers.NewInternalHandler(a.DB, env)
	internalHandler.RegisterRoutes(a.Router, constants.RouteIndex)

	plantSpeciesHandler := handlers.NewCRUDHandler(
		a.DB,
		func() *plants.PlantSpecies { return &plants.PlantSpecies{} },
		func(d *sql.DB) db.CRUDTable[plants.PlantSpecies] {
			return &db.PlantSpeciesTable{DB: d}
		},
	)
	plantSpeciesHandler.RegisterRoutes(a.Router, constants.RoutePlantSpecies)

	plantCultivarHandler := handlers.NewCRUDHandler(
		a.DB,
		func() *plants.PlantCultivar { return &plants.PlantCultivar{} },
		func(d *sql.DB) db.CRUDTable[plants.PlantCultivar] {
			return &db.PlantCultivarTable{DB: d}
		},
	)
	plantCultivarHandler.RegisterRoutes(a.Router, constants.RoutePlantCultivar)

	plantHandler := handlers.NewCRUDHandler(
		a.DB,
		func() *plants.Plant { return &plants.Plant{} },
		func(d *sql.DB) db.CRUDTable[plants.Plant] {
			return &db.PlantTable{DB: d}
		},
	)
	plantHandler.RegisterRoutes(a.Router, constants.RoutePlant)
}

// Run starts the app and now includes graceful shutdown logic.
// It uses the timeout values from your environment configuration.
func (a *App) Run(env *constants.EnvConfig) {
	// 1. Create a custom http.Server with timeouts from your config
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", env.Server.Host, env.Server.Port),
		Handler: a.Router,
		// These timeouts are critical for production-grade services.
		ReadTimeout:  env.Server.ReadTimeout,
		WriteTimeout: env.Server.WriteTimeout,
		IdleTimeout:  env.Server.IdleTimeout,
	}

	// 2. Implement Graceful Shutdown
	// This runs the server in a goroutine so that it doesn't block
	// the main thread, allowing us to listen for shutdown signals.
	go func() {
		log.Printf("INFO: Server starting on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("FATAL: Could not start server: %v\n", err)
		}
	}()

	// Create a channel to receive OS signals.
	quit := make(chan os.Signal, 1)
	// We'll listen for interrupt (Ctrl+C) and termination signals.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block execution until a signal is received.
	<-quit
	log.Println("INFO: Shutdown signal received. Starting graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), env.Server.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("FATAL: Server shutdown failed: %v", err)
	}

	log.Println("INFO: Server exited gracefully.")
}
