package main

import (
	"github.com/kylep342/mendel/internal/constants"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	env := constants.Env()

	log.Info().Str("database", env.Database.Name).Msg("Migrating db")
}
