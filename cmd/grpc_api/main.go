package main

import (
	"context"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	grpc2 "github.com/iyhunko/hash-generation-app/internal/grpc"
	pb "github.com/iyhunko/hash-generation-app/internal/proto"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	networkStr               = "tcp"
	startingServerMessage    = "Starting GRPC api server"
	failedToServeErrMessage  = "failed to serve: %s"
	failedToListenErrMessage = "failed to listen: %v"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatalf("failed to init create logger: %v.", err)
	}
	conf := config.NewConfig(lgr)
	cacheStorage := store.NewStore(lgr)

	lgr.Info(startingServerMessage)
	lis, err := net.Listen(networkStr, fmt.Sprintf(":%s", conf.GRPCServerPort))
	if err != nil {
		lgr.FatalError(fmt.Errorf(failedToListenErrMessage, err))
	}

	grpcServer := grpc.NewServer()
	hashServer := grpc2.NewHashServer(conf, cacheStorage)
	pb.RegisterHashServiceServer(grpcServer, &hashServer)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go startServer(grpcServer, lis, lgr)

	<-done
	lgr.Info("server stopped")
	shutdown(context.Background(), grpcServer, lgr)
}

func startServer(srv *grpc.Server, lis net.Listener, lgr logger.Logger) {
	if err := srv.Serve(lis); err != nil {
		lgr.FatalError(fmt.Errorf(failedToServeErrMessage, err))
	}
}

func shutdown(ctx context.Context, server *grpc.Server, lgr logger.Logger) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	server.GracefulStop()

	lgr.Info("Server shutdown done")
}
