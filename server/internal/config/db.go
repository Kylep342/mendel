package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// DatabaseConfig holds database-related configurations.
type DatabaseConfig struct {
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
}

// DSN returns the Data Source Name for connecting to the database.
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		c.Dialect, c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

// LoadDatabaseConfig loads the database-specific configuration.
// It accepts an isProd flag to perform environment-specific validation.
func LoadDatabaseConfig(v *viper.Viper, isProd bool) (*DatabaseConfig, error) {
	// Set defaults
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

	// Bind environment variables
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

	// Unmarshal the config
	var cfg DatabaseConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal database config: %w", err)
	}

	// Validate the config
	if isProd && cfg.Password == "" {
		return nil, fmt.Errorf("in production, required configuration DB_PASSWORD is not set")
	}

	return &cfg, nil
}
