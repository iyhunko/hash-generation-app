package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	http2 "github.com/iyhunko/hash-generation-app/internal/http"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	dLog "log"
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

var lgr logger.Logger
var conf config.Config
var cStorage store.Storage

func init() {
	nl, err := logger.New()
	if err != nil {
		dLog.Fatalf("failed to init create logger: %v.", err)
	}
	lgr = nl

	conf = config.NewConfig(lgr)

	cStorage = store.NewStore(lgr)
}

func main() {
	lgr.Info(startingServerMsg)
	router := http2.InitRouter(conf, cStorage)
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
