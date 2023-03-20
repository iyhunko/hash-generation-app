package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"time"
)

// Config see doc here: https://github.com/kelseyhightower/envconfig
type Config struct {
	HTTPServerPort         string        `envconfig:"HTTP_SERVER_PORT" default:"8001"`
	GRPCServerPort         string        `envconfig:"GRPC_SERVER_PORT" default:"8002"`
	WriteTimeout           time.Duration `envconfig:"WRITE_TIMEOUT" default:"15s"`
	ReadTimeout            time.Duration `envconfig:"READ_TIMEOUT" default:"15s"`
	CacheSize              int           `envconfig:"CACHE_SIZE" default:"10"`
	HashGenerationInterval time.Duration `envconfig:"HASH_GENERATION_INTERVAL" default:"5s"`
	HashKeyInCash          string        `envconfig:"HASH_KEY_IN_CACHE" default:"hash"`
}

func InitConfig() Config {
	var conf Config
	readEnv(&conf)

	return conf
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
