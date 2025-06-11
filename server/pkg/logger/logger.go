package logger

import (
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func NewLogger(deployment string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("deployment", deployment).
		Str("deployment_id", uuid.NewString()).
		Logger()
	logger.Info().Msgf("Initializing app: %s", deployment)
	return logger
}
