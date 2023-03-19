package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"time"
)

// Config see doc here: https://github.com/kelseyhightower/envconfig
type Config struct {
	HTTPServerPort string `envconfig:"HTTP_SERVER_PORT" default:"8001"`
	GRPCServerPort string `envconfig:"GRPC_SERVER_PORT" default:"8002"`

	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"15s"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"15s"`
}

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
