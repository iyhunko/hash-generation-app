package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/entity"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	dLog "log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	workerStartMsg = "Starting hash_refresher worker"
)

var lgr logger.Logger
var conf config.Config
var cStorage store.Storage

func init() {
	nl, err := logger.New()
	if err != nil {
		dLog.Fatalf("failed to init create logger: %v.", err)
	}
	lgr = nl

	conf = config.NewConfig(lgr)

	cStorage = store.NewStore(lgr)
}

func main() {
	lgr.Info(workerStartMsg)
	ticker := time.NewTicker(conf.HashGenerationInterval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				lgr.Info(fmt.Sprintf("Updated at %+v", t))
				err := refreshHash(conf, cStorage)
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
