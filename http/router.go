package http

import (
	"github.com/gorilla/mux"
	"github.com/iyhunko/hash-generation-app/config"
	"github.com/iyhunko/hash-generation-app/http/handlers"
	"github.com/iyhunko/hash-generation-app/store"
	"net/http"
)

func InitRouter(
	config config.Config,
	store store.Store,
) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	hh := handlers.NewHashHandler(config, store)
	router.HandleFunc("/hash", hh.Get).Methods(http.MethodGet)
	return router
}
