// main.go
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/kylep342/mendel/internal/app"
	"github.com/kylep342/mendel/internal/constants"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	env := constants.Env()

	log.Info().Msg("Initializing app")
	a := app.App{}
	a.Initialize(env)

	log.Info().Msg("App initialized. Running")
	a.Run(env)
}
