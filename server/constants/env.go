package constants

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// EnvConfig is a singleton struct containing runtime constants
// it holds all environment-specific configurations for the application.
// Struct tags define how environment variables are mapped to fields.
type EnvConfig struct {
	// Server configuration fields
	Server struct {
		Host            string        `envconfig:"HOST" default:"localhost"`       //  SERVER_HOST
		Port            string        `envconfig:"PORT" default:"8080"`            //  SERVER_PORT
		ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT" default:"15s"`     //  SERVER_READ_TIMEOUT
		WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT" default:"15s"`    //  SERVER_WRITE_TIMEOUT
		IdleTimeout     time.Duration `envconfig:"IDLE_TIMEOUT" default:"60s"`     //  SERVER_IDLE_TIMEOUT
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"` //  SERVER_SHUTDOWN_TIMEOUT
	} `envconfig:"SERVER"` // prefix members with "SERVER_"

	// Database configuration fields
	Database struct {
		Dialect         string        `envconfig:"DIALECT" default:"postgres"`
		Host            string        `envconfig:"HOST" default:"localhost"`       //  DB_HOST
		Port            int           `envconfig:"PORT" default:"5432"`            //  DB_PORT
		User            string        `envconfig:"USER" default:"postgres"`        //  DB_USER
		Password        string        `envconfig:"PASSWORD" required:"false"`      //  DB_PASSWORD (consider required:"true")
		Name            string        `envconfig:"NAME" default:"mendel_db"`       //  DB_NAME
		SSLMode         string        `envconfig:"SSLMODE" default:"disable"`      //  DB_SSLMODE
		MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"25"`    //  DB_MAX_OPEN_CONNS
		MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"25"`    //  DB_MAX_IDLE_CONNS
		ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"5m"` //  DB_CONN_MAX_LIFETIME
	} `envconfig:"DB"` // prefix members with "DB_"

	// Application specific configuration
	App struct {
		Name                string `envconfig:"NAME" default:"MendelApp"`                  //  APP_NAME
		Environment         string `envconfig:"ENV" default:"development" required:"true"` //  APP_ENV
		LogLevel            string `envconfig:"LOG_LEVEL" default:"info"`                  //  APP_LOG_LEVEL
		EnableDebugFeatures bool   `envconfig:"ENABLE_DEBUG_FEATURES" default:"false"`     //  APP_ENABLE_DEBUG_FEATURES
	} `envconfig:"APP"` // prefix members with "APP_"
}

// globalEnvConfig holds the loaded environment configuration (singleton).
var globalEnvConfig *EnvConfig

// LoadEnv loads configuration from environment variables into the EnvConfig struct
// using the `envconfig` library which processes struct tags.
// It should be called once at application startup.
func LoadEnv() *EnvConfig {
	if globalEnvConfig != nil {
		return globalEnvConfig // Already loaded
	}

	var cfg EnvConfig
	// The first argument to Process is a global prefix for all environment variables.
	// Since we are using prefixes on the nested struct fields (e.g., `envconfig:"SERVER"`),
	// we pass an empty string here. The library will combine these prefixes.
	// For example, `envconfig:"SERVER"` and `envconfig:"PORT"` results in `SERVER_PORT`.
	err := envconfig.Process("", &cfg)
	if err != nil {
		// The library provides detailed errors, e.g., for missing required fields
		// or type conversion errors.
		log.Fatalf("FATAL: Failed to process environment configuration: %v", err)
	}

	// Optional: Additional custom validation after processing
	if cfg.App.Environment == "production" && cfg.Database.Password == "" {
		// Note: `required:"true"` on the Password field for production is often preferred.
		// This is a secondary check or if `required` isn't granular enough.
		log.Printf("WARN: DB_PASSWORD environment variable is not set in production. This might be a security risk or cause connection failure.")
	}
	// You can add more complex cross-field validations here if needed.

	globalEnvConfig = &cfg
	log.Println("INFO: Environment configuration loaded successfully.")
	return globalEnvConfig
}

// GetEnv returns the loaded environment configuration.
// It panics if LoadEnv has not been called first.
func GetEnv() *EnvConfig {
	if globalEnvConfig == nil {
		log.Fatal("FATAL: Environment configuration has not been loaded. Call constants.LoadEnv() in main.")
	}
	return globalEnvConfig
}
