package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/entity"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"log"
	"time"
)

const (
	workerStartMsg = "Starting hash_refresher worker"
	updatedMsg     = "Updated: %s"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	lgr.Info(workerStartMsg)
	conf := config.NewConfig(lgr)
	storage := store.NewStore(lgr)

	for range time.Tick(conf.HashGenerationInterval) {
		err = refreshHash(lgr, conf, storage)
		if err != nil {
			lgr.ErrorWithExit(err.Error())
		}
	}
}

func refreshHash(lgr logger.Logger, conf config.Config, storage store.Storage) error {
	hash := entity.NewHash()
	marshaledHash, err := json.Marshal(hash)
	if err != nil {
		return err
	}
	err = storage.Set(conf.HashFilePath, marshaledHash)
	if err != nil {
		return err
	}
	lgr.Info(fmt.Sprintf(updatedMsg, hash))
	return nil
}
