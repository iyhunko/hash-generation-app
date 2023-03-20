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

const (
	startingServerMsg  = "Starting http api server"
	listeningToPortMsg = "Listening to %s port..."
	listeningErrMsg    = "Error listening %s"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	lgr.Info(startingServerMsg)
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

	lgr.Info(fmt.Sprintf(listeningToPortMsg, conf.HTTPServerPort))
	lgr.ErrorWithExit(fmt.Sprintf(listeningErrMsg, srv.ListenAndServe()))
}
