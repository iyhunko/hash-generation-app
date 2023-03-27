package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	http2 "github.com/iyhunko/hash-generation-app/internal/http"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	startingServerMsg = "Starting http api server"
	listeningErrMsg   = "error listening %w"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatalf("failed to init create logger: %v.", err)
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

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go startServer(srv, lgr)

	<-done
	lgr.Info("server stopped")
	shutdown(context.Background(), srv, lgr)
}

func startServer(srv *http.Server, lgr logger.Logger) {
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		lgr.FatalError(fmt.Errorf(listeningErrMsg, err))
	}
}

func shutdown(ctx context.Context, server *http.Server, lgr logger.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		lgr.FatalError(fmt.Errorf("server shutdown failed: %w", err))
	}
	lgr.Info("Server shutdown done")
}
