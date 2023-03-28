package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/entity"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	workerStartMsg = "Starting hash_refresher worker"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatalf("failed to init create logger: %v.", err)
	}

	lgr.Info(workerStartMsg)
	conf := config.NewConfig(lgr)
	storage := store.NewStore(lgr)

	ticker := time.NewTicker(conf.HashGenerationInterval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				lgr.Info(fmt.Sprintf("Updated at %+v", t))
				err = refreshHash(conf, storage)
				if err != nil {
					lgr.FatalError(err)
				}
			}
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Wait for interrupt signal to gracefully shutdown the ticker
	ticker.Stop()
	done <- true
}

func refreshHash(conf config.Config, storage store.Storage) error {
	hash := entity.NewHash()
	marshaledHash, err := json.Marshal(hash)
	if err != nil {
		return fmt.Errorf("failed to marshal hash: %w", err)
	}
	err = storage.Set(conf.HashFilePath, marshaledHash)
	if err != nil {
		return fmt.Errorf("failed to set value to storage: %w", err)
	}
	return nil
}
