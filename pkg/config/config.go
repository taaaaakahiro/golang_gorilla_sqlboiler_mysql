package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port int `env:"PORT,required"`
	DB   *databaseConfig
}

type databaseConfig struct {
	Dsn string `env:"MYSQL_DSN,required"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
