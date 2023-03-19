package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/iyhunko/hash-generation-app/config"
	"log"
	"net/http"
	"time"
)

func hashHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))

	id := uuid.New()
	w.Write([]byte("\nGenerated UUID: " + id.String()))
}

func main() {
	log.Println("Starting http api server")

	// load env variables to the Config struct
	var conf config.Config
	config.ReadEnv(&conf)

	// init http server
	mux := http.NewServeMux()
	mux.HandleFunc("/hash", hashHandler)
	srv := &http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf(":%s", conf.HTTPServerPort),
		WriteTimeout: conf.WriteTimeout,
		ReadTimeout:  conf.ReadTimeout,
	}

	log.Printf("Listening to %s port...", conf.HTTPServerPort)
	log.Fatal(context.Background(), "error listening ", srv.ListenAndServe())
}
