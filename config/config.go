package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Auth `yaml:"auth"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port            string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		ReadTimeout     int    `env-required:"true" yaml:"read_timeout" env:"READ_TIMEOUT"`
		WriteTimeout    int    `env-required:"true" yaml:"write_timeout" env:"WRITE_TIMEOUT"`
		ShutdownTimeout int    `env-required:"true" yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT"`
	}

	// Auth -.
	Auth struct {
		SigningKey string `env-required:"true" yaml:"signing_key" env:"SIGNING_KEY"`
		HashSalt   string `env-required:"true" yaml:"hash_salt" env:"HASH_SALT"`
		TokenTTL   int    `env-required:"true" yaml:"token_ttl" env:"TOKEN_TTL"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}
)

// Возвращает инстанс конфига
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
