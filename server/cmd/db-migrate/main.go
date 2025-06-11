package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/kylep342/mendel/internal/constants"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("deployment", constants.AppDbMigrate).
		Str("deployment_id", uuid.NewString()).
		Logger()

	env := constants.Env(logger)

	logger.Info().Str("database", env.Database.Name).Msg("Migrating database")
	m, err := migrate.New(
		"file://"+env.Database.MigrationsFolder,
		env.DBUrl(),
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create a database migration instance")
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		version, _, _ := m.Version()
		logger.Info().Uint("migration", version).Msg("No migration necessary")
		return
	} else if err != nil {
		logger.Fatal().Err(err).Msg("Failed to migrate database")
	}

	logger.Info().Str("database", env.Database.Name).Msg("Database migrated")
}
