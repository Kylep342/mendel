package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/kylep342/mendel/internal/constants"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	env := constants.Env()

	log.Info().Str("database", env.Database.Name).Msg("Migrating database")

	m, err := migrate.New(
		"file://"+env.Database.MigrationsFolder,
		env.DBUrl(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create a database migration instance")
	}

	err = m.Up()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}
	log.Info().Str("database", env.Database.Name).Msg("Database migrated")
}
