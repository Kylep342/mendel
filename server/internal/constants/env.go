package constants

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

// EnvConfig is a singleton struct containing runtime constants
// it holds all environment-specific configurations for the application.
// Struct tags define how environment variables are mapped to fields.
type EnvConfig struct {
	// Server configuration fields
	Server struct {
		Host            string        `json:"host" envconfig:"HOST" default:"localhost"`                   //  SERVER_HOST
		Port            string        `json:"port" envconfig:"PORT" default:"8080"`                        //  SERVER_PORT
		ReadTimeout     time.Duration `json:"read_timeout" envconfig:"READ_TIMEOUT" default:"15s"`         //  SERVER_READ_TIMEOUT
		WriteTimeout    time.Duration `json:"write_timeout" envconfig:"WRITE_TIMEOUT" default:"15s"`       //  SERVER_WRITE_TIMEOUT
		IdleTimeout     time.Duration `json:"idle_timeout" envconfig:"IDLE_TIMEOUT" default:"60s"`         //  SERVER_IDLE_TIMEOUT
		ShutdownTimeout time.Duration `json:"shutdown_timeout" envconfig:"SHUTDOWN_TIMEOUT" default:"10s"` //  SERVER_SHUTDOWN_TIMEOUT
	} `json:"server" envconfig:"SERVER"` // prefix members with "SERVER_"

	// Database configuration fields
	Database struct {
		Dialect          string        `json:"dialect" envconfig:"DIALECT" default:"postgres"`                                        //  DB_DIALECT
		Host             string        `json:"host" envconfig:"HOST" default:"localhost"`                                             //  DB_HOST
		Port             int           `json:"port" envconfig:"PORT" default:"5432"`                                                  //  DB_PORT
		User             string        `json:"user" envconfig:"USER" default:"postgres"`                                              //  DB_USER
		Password         string        `json:"password" envconfig:"PASSWORD" required:"true"`                                         //  DB_PASSWORD (consider required:"true")
		Name             string        `json:"name" envconfig:"NAME" default:"mendel_db"`                                             //  DB_NAME
		SSLMode          string        `json:"ssl_mode" envconfig:"SSLMODE" default:"disable"`                                        //  DB_SSLMODE
		MaxOpenConns     int           `json:"max_openconns" envconfig:"MAX_OPEN_CONNS" default:"25"`                                 //  DB_MAX_OPEN_CONNS
		MaxIdleConns     int           `json:"max_idle_conns" envconfig:"MAX_IDLE_CONNS" default:"25"`                                //  DB_MAX_IDLE_CONNS
		ConnMaxLifetime  time.Duration `json:"conn_max_lifetime" envconfig:"CONN_MAX_LIFETIME" default:"5m"`                          //  DB_CONN_MAX_LIFETIME
		MigrationsFolder string        `json:"migrations_folder" envconfig:"MIGRATIONS_FOLDER" default:"/app/internal/db/migrations"` // DB_MIGRATIONS_FOLDER
	} `json:"database" envconfig:"DB"` // prefix members with "DB_"

	// Application specific configuration
	App struct {
		Name                string `json:"name" envconfig:"NAME" default:"MendelApp"`                               //  APP_NAME
		Environment         string `json:"environment" envconfig:"ENV" default:"development" required:"true"`       //  APP_ENV
		LogLevel            string `json:"log_level" envconfig:"LOG_LEVEL" default:"info"`                          //  APP_LOG_LEVEL
		EnableDebugFeatures bool   `json:"enable_debug_features" envconfig:"ENABLE_DEBUG_FEATURES" default:"false"` //  APP_ENABLE_DEBUG_FEATURES
	} `json:"app" envconfig:"APP"` // prefix members with "APP_"
}

func (e *EnvConfig) DBUrl() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		e.Database.Dialect,
		e.Database.User,
		e.Database.Password,
		e.Database.Host,
		e.Database.Port,
		e.Database.Name,
		e.Database.SSLMode,
	)
}

// globalEnvConfig holds the loaded environment configuration (singleton).
var (
	globalEnvConfig *EnvConfig
	loadConfigOnce  sync.Once

	allowedEnvironments = []string{EnvDevelopment, EnvStaging, EnvProduction}
	// allowedLogLevels    = []zerolog.Level{
	// 	zerolog.DebugLevel,
	// 	zerolog.InfoLevel,
	// 	zerolog.WarnLevel,
	// 	zerolog.ErrorLevel,
	// 	zerolog.FatalLevel,
	// 	zerolog.PanicLevel,
	// 	zerolog.NoLevel,
	// 	zerolog.Disabled,
	// 	zerolog.TraceLevel,
	// }
)

func isValidValue(value string, allowedValues []string, caseSensitive bool) bool {
	for _, allowed := range allowedValues {
		if caseSensitive {
			if value == allowed {
				return true
			}
		} else {
			if strings.EqualFold(value, allowed) { // Case-insensitive comparison
				return true
			}
		}
	}
	return false
}

// loadEnv contains the logic to load and validate environment configuration.
func loadEnv(logger zerolog.Logger) {
	logger.Info().Msg("Initializing and loading environment configuration")
	var cfg EnvConfig
	envconfig.MustProcess("", &cfg)

	// Validate environment
	if !isValidValue(cfg.App.Environment, allowedEnvironments, true) {
		logger.Fatal().Msgf("Invalid APP_ENV value '%s'. Allowed values are: %v",
			cfg.App.Environment, allowedEnvironments)
	}

	if cfg.App.Environment == EnvProduction && cfg.Database.Password == "" {
		logger.Warn().Msg("DB_PASSWORD environment variable is not set in production. This might be a security risk or cause connection failure")
	}

	globalEnvConfig = &cfg
	logger.Info().Msg("Environment configuration loaded successfully")
}

// Env returns the loaded environment configuration.
// It ensures that the configuration is loaded exactly once, in a thread-safe manner.
func Env(logger zerolog.Logger) *EnvConfig {
	loadConfigOnce.Do(func() { loadEnv(logger) })

	if globalEnvConfig == nil {
		logger.Fatal().Msg("Environment configuration is nil after attempting to load")
	}
	return globalEnvConfig
}
