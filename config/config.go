package config

import (
	"time"

	"github.com/pkg/errors"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Api struct {
		Host        string        `envconfig:"-"`
		Port        string        `envconfig:"default=8880"`
		ReadTimeout time.Duration `envconfig:"default=10s"`
	}
	Log struct {
		Format string `envconfig:"default=text"`
		Level  string `envconfig:"default=debug"`
	}
	Postgres struct {
		Dsn string `envconfig:"default=postgres://postgres:postgres@localhost:5432/g3?sslmode=disable&connect_timeout=3&binary_parameters=yes"`
	}
}

func Load() (*Config, error) {
	var conf *Config
	if err := envconfig.Init(&conf); err != nil {
		return nil, errors.Wrap(err, "config: can't load config")
	}
	return conf, nil
}
