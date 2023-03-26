package main

import (
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	grpc2 "github.com/iyhunko/hash-generation-app/internal/grpc"
	pb "github.com/iyhunko/hash-generation-app/internal/proto"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"google.golang.org/grpc"
	"log"
	"net"
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

	lgr.Info(startingServerMessage)
	conf := config.NewConfig(lgr)
	cacheStorage := store.NewStore(lgr)

	lis, err := net.Listen(networkStr, fmt.Sprintf(":%s", conf.GRPCServerPort))
	if err != nil {
		lgr.FatalError(fmt.Errorf(failedToListenErrMessage, err))
	}

	grpcServer := grpc.NewServer()
	hashServer := grpc2.NewHashServer(conf, cacheStorage)
	pb.RegisterHashServiceServer(grpcServer, &hashServer)

	if err := grpcServer.Serve(lis); err != nil {
		lgr.FatalError(fmt.Errorf(failedToServeErrMessage, err))
	}
}
