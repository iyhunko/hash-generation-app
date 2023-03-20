package main

import (
	"context"
	"fmt"
	"github.com/iyhunko/hash-generation-app/config"
	http2 "github.com/iyhunko/hash-generation-app/http"
	"github.com/iyhunko/hash-generation-app/store"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting http api server")

	conf := config.InitConfig()
	cacheStorage := store.NewStore()

	// init http server and router
	router := http2.InitRouter(conf, cacheStorage)
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", conf.HTTPServerPort),
		WriteTimeout: conf.WriteTimeout,
		ReadTimeout:  conf.ReadTimeout,
	}

	log.Printf("Listening to %s port...", conf.HTTPServerPort)
	log.Fatal(context.Background(), "error listening ", srv.ListenAndServe())
}
