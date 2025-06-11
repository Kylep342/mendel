// main.go
package main

import (
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/kylep342/mendel/internal/app"
	"github.com/kylep342/mendel/internal/constants"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("deployment", constants.AppMendelServer).
		Str("deployment_id", uuid.NewString()).
		Logger()

	env := constants.Env(logger)

	logger.Info().Msg("Initializing app")
	a := app.App{}
	a.Initialize(logger, env)

	logger.Info().Msg("App initialized. Running")
	a.Run(env)
}
