package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// ServerConfig holds server-related configurations.
type ServerConfig struct {
	Host            string        `json:"host" mapstructure:"host"`
	Port            string        `json:"port" mapstructure:"port"`
	ReadTimeout     time.Duration `json:"read_timeout" mapstructure:"readtimeout"`
	WriteTimeout    time.Duration `json:"write_timeout" mapstructure:"writetimeout"`
	IdleTimeout     time.Duration `json:"idle_timeout" mapstructure:"idletimeout"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout" mapstructure:"shutdowntimeout"`
}

// Addr returns the full server address (host:port).
func (c *ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// LoadServerConfig loads the server-specific configuration.
func LoadServerConfig(v *viper.Viper) (*ServerConfig, error) {
	// Set defaults
	v.SetDefault("server.host", "localhost")
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.readtimeout", "15s")
	v.SetDefault("server.writetimeout", "15s")
	v.SetDefault("server.idletimeout", "60s")
	v.SetDefault("server.shutdowntimeout", "10s")

	// Bind environment variables
	v.BindEnv("server.host", "SERVER_HOST")
	v.BindEnv("server.port", "SERVER_PORT")
	v.BindEnv("server.readtimeout", "SERVER_READ_TIMEOUT")
	v.BindEnv("server.writetimeout", "SERVER_WRITE_TIMEOUT")
	v.BindEnv("server.idletimeout", "SERVER_IDLE_TIMEOUT")
	v.BindEnv("server.shutdowntimeout", "SERVER_SHUTDOWN_TIMEOUT")

	// Unmarshal and return
	var cfg ServerConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal server config: %w", err)
	}

	// No specific validation needed for server config in this example,
	// but it could be added here.

	return &cfg, nil
}
