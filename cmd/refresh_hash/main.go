package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyhunko/hash-generation-app/config"
	"github.com/iyhunko/hash-generation-app/entity"
	"github.com/iyhunko/hash-generation-app/store"
	"log"
	"time"
)

func main() {
	log.Println("Starting refresh_hash worker")
	conf := config.InitConfig()
	cStorage := store.NewStore(conf.CacheSize, conf.HashGenerationInterval)

	for range time.Tick(conf.HashGenerationInterval) {
		hash := entity.NewHash()
		marshaledHash, err := json.Marshal(hash)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = cStorage.Set(conf.HashKeyInCash, marshaledHash)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Updated:", hash)
	}
}
