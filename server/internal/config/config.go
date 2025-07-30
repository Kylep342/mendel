package config

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// EnvConfig is the root configuration struct that composes all other config modules.
type EnvConfig struct {
	App      *AppConfig
	Server   *ServerConfig
	Database *DatabaseConfig
}

var (
	globalEnvConfig *EnvConfig
	loadConfigOnce  sync.Once
)

// loadConfig orchestrates the loading of all configuration modules.
func loadConfig(logger zerolog.Logger) {
	logger.Info().Msg("Initializing and loading environment configuration")
	v := viper.New()

	var cfg EnvConfig
	var err error

	// Load App config first, as others might depend on it (e.g., for env)
	cfg.App, err = LoadAppConfig(v)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load application configuration")
	}

	// Load Server config
	cfg.Server, err = LoadServerConfig(v)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load server configuration")
	}

	// Load Database config, passing the production flag from the App config
	cfg.Database, err = LoadDatabaseConfig(v, cfg.App.IsProduction())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load database configuration")
	}

	globalEnvConfig = &cfg
	logger.Info().Str("environment", cfg.App.Environment).Msg("Environment configuration loaded successfully")
}

// Env provides global, thread-safe access to the loaded configuration.
func Env(logger zerolog.Logger) *EnvConfig {
	loadConfigOnce.Do(func() {
		loadConfig(logger)
	})

	if globalEnvConfig == nil {
		logger.Fatal().Msg("Environment configuration is nil after attempting to load")
	}
	return globalEnvConfig
}
