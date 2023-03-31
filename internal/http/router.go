package http

import (
	"github.com/gorilla/mux"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/http/handler"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"net/http"
)

func InitRouter(
	config config.Config,
	store store.Storage,
) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	hh := handler.NewHashHandler(config, store)
	router.HandleFunc("/hash", hh.Get).Methods(http.MethodGet)
	return router
}
