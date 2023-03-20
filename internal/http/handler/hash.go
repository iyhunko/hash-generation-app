package handler

import (
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"net/http"
)

const (
	failedToFetchMsg = "Failed to fetch hash"
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
	if hashBytes == nil {
		http.Error(w, failedToFetchMsg, http.StatusInternalServerError)
		return
	}

	_, err := w.Write(hashBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
