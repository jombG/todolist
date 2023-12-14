package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
	"time"
)

type (
	Config struct {
		Server ServerConfig
		DB     DBConfig
	}

	ServerConfig struct {
		LogLevel zapcore.Level `env:"APP_LOG_LEVEL" envDefault:"DEBUG"`
		AppENV   string        `env:"APP_ENV" envDefault:"local"`
		AppName  string        `env:"APP_NAME" envDefault:"unknown"`

		Port              int           `env:"APP_SERVER_PORT" envDefault:"80"`
		ShutdownTimeout   time.Duration `env:"APP_SHUTDOWN_TIMEOUT" envDefault:"15s"`
		ReadHeaderTimeout time.Duration `env:"APP_SERVER_READ_HEADER_TIMEOUT" envDefault:"15s"`
	}

	DBConfig struct {
		Driver    string `env:"APP_DB_DRIVER" envDefault:""`
		DSN       string `env:"APP_DB_DSN" envDefault:""`
		Migration string `env:"APP_DB_MIGRATION" envDefault:""`
	}
)

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env.local")

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.Wrap(err, "init env config")
	}

	return &cfg, nil
}

func (c DBConfig) DataSourceName() string {
	return c.DSN
}

func (c DBConfig) DriverName() string {
	return c.Driver
}

func (c DBConfig) MigrationSource() string {
	return c.Migration
}
