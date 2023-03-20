package main

import (
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	http2 "github.com/iyhunko/hash-generation-app/internal/http"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"log"
	"net/http"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	lgr.Info("Starting http api server")
	conf := config.NewConfig(lgr)
	cacheStorage := store.NewStore(lgr)

	// init http server and router
	router := http2.InitRouter(conf, cacheStorage)
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", conf.HTTPServerPort),
		WriteTimeout: conf.WriteTimeout,
		ReadTimeout:  conf.ReadTimeout,
	}

	lgr.Info(fmt.Sprintf("Listening to %s port...", conf.HTTPServerPort))
	lgr.ErrorWithExit(fmt.Sprintf("Error listening %s ", srv.ListenAndServe()))
}
