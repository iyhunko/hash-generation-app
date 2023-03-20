package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyhunko/hash-generation-app/config"
	"github.com/iyhunko/hash-generation-app/entity"
	"github.com/iyhunko/hash-generation-app/logger"
	"github.com/iyhunko/hash-generation-app/store"
	"log"
	"time"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	lgr.Info("Starting refresh_hash worker")
	conf := config.NewConfig(lgr)
	storage := store.NewStore(lgr)

	for range time.Tick(conf.HashGenerationInterval) {
		hash := entity.NewHash()
		marshaledHash, err := json.Marshal(hash)
		if err != nil {
			lgr.ErrorWithExit(err.Error())
		}
		err = storage.Set(conf.HashFilePath, marshaledHash)
		if err != nil {
			lgr.ErrorWithExit(err.Error())
		}
		lgr.Info(fmt.Sprintf("Updated: %s", hash))
	}
}
