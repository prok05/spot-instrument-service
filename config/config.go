package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App        App
		Log        Log
		GRPC       GRPC
		Jaeger     Jaeger
		Prometheus Prometheus
	}

	App struct {
		Name string `env:"APP_NAME,required"`
		ENV  string `env:"APP_ENVIRONMENT,required"`
	}

	GRPC struct {
		Port string `env:"GRPC_PORT,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	Jaeger struct {
		GrpcAddr string `env:"JAEGER_GRPC_ADDR,required"`
	}

	Prometheus struct {
		Addr string `env:"PROMETHEUS_ADDR,required"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
