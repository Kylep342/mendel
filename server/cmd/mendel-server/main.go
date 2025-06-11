// main.go
package main

import (
	"github.com/kylep342/mendel/internal/app"
	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/pkg/logger"
)

func main() {
	logger := logger.NewLogger(constants.AppMendelServer)
	env := constants.Env(logger)

	a := app.App{}
	a.Initialize(logger, env)

	logger.Info().Msg("App initialized. Running")
	a.Run(env)
}
