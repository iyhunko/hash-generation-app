package handlers

import (
	"encoding/json"
	"github.com/iyhunko/hash-generation-app/config"
	"github.com/iyhunko/hash-generation-app/entity"
	"github.com/iyhunko/hash-generation-app/store"
	"log"
	"net/http"
)

type HashHandler struct {
	config config.Config
	store  store.Store
}

func NewHashHandler(
	config config.Config,
	store store.Store,
) HashHandler {
	return HashHandler{
		config: config,
		store:  store,
	}
}

func (hh *HashHandler) Get(w http.ResponseWriter, r *http.Request) {
	hashBytes := hh.store.Get(hh.config.HashKeyInCash)
	if hashBytes == nil {
		hash := entity.NewHash()
		marshaledHash, err := json.Marshal(hash)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = hh.store.Set(hh.config.HashKeyInCash, marshaledHash)
		if err != nil {
			log.Fatal(err.Error())
		}
		hashBytes = marshaledHash
	}

	_, err := w.Write(hashBytes)
	if err != nil {
		log.Fatal("Failed to send response")
	}
	w.WriteHeader(http.StatusOK)
	return
}
