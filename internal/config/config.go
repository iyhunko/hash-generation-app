package config

import (
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"github.com/kelseyhightower/envconfig"
	"time"
)

// Config see doc here: https://github.com/kelseyhightower/envconfig
type Config struct {
	HTTPServerPort         string        `envconfig:"HTTP_SERVER_PORT" default:"8001"`
	GRPCServerPort         string        `envconfig:"GRPC_SERVER_PORT" default:"8002"`
	WriteTimeout           time.Duration `envconfig:"WRITE_TIMEOUT" default:"15s"`
	ReadTimeout            time.Duration `envconfig:"READ_TIMEOUT" default:"15s"`
	HashGenerationInterval time.Duration `envconfig:"HASH_GENERATION_INTERVAL" default:"5s"`
	HashFilePath           string        `envconfig:"HASH_FILE_PATH" default:"hash.json"`
}

func NewConfig(log logger.Logger) Config {
	var conf Config
	// TODO: refactor
	err := envconfig.Process("", &conf)
	if err != nil {
		log.FatalError(err)
	}

	return conf
}
