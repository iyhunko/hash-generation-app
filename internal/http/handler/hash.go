package handler

import (
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/store"
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
	hashBytes := hh.store.Get(hh.config.HashFilePath)

	_, err := w.Write(hashBytes)
	if err != nil {
		log.Fatal("Failed to send response")
	}
	return
}
