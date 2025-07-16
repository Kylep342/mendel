package constants

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// EnvConfig is a singleton struct containing runtime constants.
// It holds all environment-specific configurations for the application.
// The `mapstructure` tags are used by Viper to unmarshal config data into this struct.
type EnvConfig struct {
	// Server configuration fields
	Server struct {
		Host            string        `json:"host" mapstructure:"host"`
		Port            string        `json:"port" mapstructure:"port"`
		ReadTimeout     time.Duration `json:"read_timeout" mapstructure:"readtimeout"`
		WriteTimeout    time.Duration `json:"write_timeout" mapstructure:"writetimeout"`
		IdleTimeout     time.Duration `json:"idle_timeout" mapstructure:"idletimeout"`
		ShutdownTimeout time.Duration `json:"shutdown_timeout" mapstructure:"shutdowntimeout"`
	} `json:"server" mapstructure:"server"`

	// Database configuration fields
	Database struct {
		Dialect          string        `json:"dialect" mapstructure:"dialect"`
		Host             string        `json:"host" mapstructure:"host"`
		Port             int           `json:"port" mapstructure:"port"`
		User             string        `json:"user" mapstructure:"user"`
		Password         string        `json:"password" mapstructure:"password"`
		Name             string        `json:"name" mapstructure:"name"`
		SSLMode          string        `json:"ssl_mode" mapstructure:"sslmode"`
		MaxOpenConns     int           `json:"max_openconns" mapstructure:"maxopenconns"`
		MaxIdleConns     int           `json:"max_idle_conns" mapstructure:"maxidleconns"`
		ConnMaxLifetime  time.Duration `json:"conn_max_lifetime" mapstructure:"connmaxlifetime"`
		MigrationsFolder string        `json:"migrations_folder" mapstructure:"migrationsfolder"`
	} `json:"database" mapstructure:"database"`

	// Application specific configuration
	App struct {
		Name                string `json:"name" mapstructure:"name"`
		Environment         string `json:"environment" mapstructure:"environment"`
		LogLevel            string `json:"log_level" mapstructure:"loglevel"`
		EnableDebugFeatures bool   `json:"enable_debug_features" mapstructure:"enabledebugfeatures"`
		WebHost             string `json:"web_host" mapstructure:"webhost"`
	} `json:"app" mapstructure:"app"`
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

func loadEnv(logger zerolog.Logger) {
	logger.Info().Msg("Initializing and loading environment configuration with Viper")
	v := viper.New()

	v.SetDefault("server.host", "localhost")
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.readtimeout", "15s")
	v.SetDefault("server.writetimeout", "15s")
	v.SetDefault("server.idletimeout", "60s")
	v.SetDefault("server.shutdowntimeout", "10s")

	v.SetDefault("database.dialect", "postgres")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.name", "mendel_db")
	v.SetDefault("database.sslmode", "disable")
	v.SetDefault("database.maxopenconns", 25)
	v.SetDefault("database.maxidleconns", 25)
	v.SetDefault("database.connmaxlifetime", "5m")
	v.SetDefault("database.migrationsfolder", "/app/internal/db/migrations")

	v.SetDefault("app.name", "MendelApp")
	v.SetDefault("app.environment", "development")
	v.SetDefault("app.loglevel", "info")
	v.SetDefault("app.enabledebugfeatures", false)
	v.SetDefault("app.webhost", "http://localhost:5173")

	v.BindEnv("server.host", "SERVER_HOST")
	v.BindEnv("server.port", "SERVER_PORT")
	v.BindEnv("server.readtimeout", "SERVER_READ_TIMEOUT")
	v.BindEnv("server.writetimeout", "SERVER_WRITE_TIMEOUT")
	v.BindEnv("server.idletimeout", "SERVER_IDLE_TIMEOUT")
	v.BindEnv("server.shutdowntimeout", "SERVER_SHUTDOWN_TIMEOUT")

	v.BindEnv("database.dialect", "DB_DIALECT")
	v.BindEnv("database.host", "DB_HOST")
	v.BindEnv("database.port", "DB_PORT")
	v.BindEnv("database.user", "DB_USER")
	v.BindEnv("database.password", "DB_PASSWORD")
	v.BindEnv("database.name", "DB_NAME")
	v.BindEnv("database.sslmode", "DB_SSLMODE")
	v.BindEnv("database.maxopenconns", "DB_MAX_OPEN_CONNS")
	v.BindEnv("database.maxidleconns", "DB_MAX_IDLE_CONNS")
	v.BindEnv("database.connmaxlifetime", "DB_CONN_MAX_LIFETIME")
	v.BindEnv("database.migrationsfolder", "DB_MIGRATIONS_FOLDER")

	v.BindEnv("app.name", "APP_NAME")
	v.BindEnv("app.environment", "APP_ENV")
	v.BindEnv("app.loglevel", "APP_LOG_LEVEL")
	v.BindEnv("app.enabledebugfeatures", "APP_ENABLE_DEBUG_FEATURES")
	v.BindEnv("app.webhost", "APP_WEB_HOST")

	var cfg EnvConfig
	if err := v.Unmarshal(&cfg); err != nil {
		logger.Fatal().Err(err).Msg("Failed to unmarshal configuration")
	}

	if cfg.App.Environment == EnvProduction && cfg.Database.Password == "" {
		logger.Fatal().Msg("Required configuration DB_PASSWORD is not set")
	}

	if !isValidValue(cfg.App.Environment, allowedEnvironments, true) {
		logger.Fatal().Msgf("Invalid APP_ENV value '%s'. Allowed values are: %v",
			cfg.App.Environment, allowedEnvironments)
	}

	globalEnvConfig = &cfg
	logger.Info().Msg("Environment configuration loaded successfully")
}

// Env returns the loaded environment configuration.
func Env(logger zerolog.Logger) *EnvConfig {
	loadConfigOnce.Do(func() { loadEnv(logger) })

	if globalEnvConfig == nil {
		logger.Fatal().Msg("Environment configuration is nil after attempting to load")
	}
	return globalEnvConfig
}
