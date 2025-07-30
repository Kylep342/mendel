package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Environment constants
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

var allowedEnvironments = []string{EnvDevelopment, EnvStaging, EnvProduction}

// AppConfig holds application-specific configurations.
type AppConfig struct {
	Name                string `json:"name" mapstructure:"name"`
	Environment         string `json:"environment" mapstructure:"environment"`
	LogLevel            string `json:"log_level" mapstructure:"loglevel"`
	EnableDebugFeatures bool   `json:"enable_debug_features" mapstructure:"enabledebugfeatures"`
	WebHost             string `json:"web_host" mapstructure:"webhost"`
}

// IsProduction returns true if the app environment is set to production.
func (c *AppConfig) IsProduction() bool {
	return c.Environment == EnvProduction
}

// LoadAppConfig loads the app-specific configuration.
func LoadAppConfig(v *viper.Viper) (*AppConfig, error) {
	// Set defaults
	v.SetDefault("app.name", "MendelApp")
	v.SetDefault("app.environment", EnvDevelopment)
	v.SetDefault("app.loglevel", "info")
	v.SetDefault("app.enabledebugfeatures", false)
	v.SetDefault("app.webhost", "http://localhost:5173")

	// Bind environment variables
	v.BindEnv("app.name", "APP_NAME")
	v.BindEnv("app.environment", "APP_ENV")
	v.BindEnv("app.loglevel", "APP_LOG_LEVEL")
	v.BindEnv("app.enabledebugfeatures", "APP_ENABLE_DEBUG_FEATURES")
	v.BindEnv("app.webhost", "APP_WEB_HOST")

	// Unmarshal the config
	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal app config: %w", err)
	}

	// Validate the config
	if !isValidValue(cfg.Environment, allowedEnvironments, true) {
		return nil, fmt.Errorf("invalid APP_ENV value '%s'. Allowed values are: %v",
			cfg.Environment, allowedEnvironments)
	}

	return &cfg, nil
}

// isValidValue checks if a value is in a list of allowed values.
func isValidValue(value string, allowedValues []string, caseSensitive bool) bool {
	for _, allowed := range allowedValues {
		if caseSensitive {
			if value == allowed {
				return true
			}
		} else {
			if strings.EqualFold(value, allowed) {
				return true
			}
		}
	}
	return false
}
