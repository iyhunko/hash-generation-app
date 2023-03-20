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

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	lgr.Info("Starting GRPC api server")
	conf := config.NewConfig(lgr)
	cacheStorage := store.NewStore(lgr)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCServerPort))
	if err != nil {
		lgr.ErrorWithExit(fmt.Sprintf("failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()
	hashServer := grpc2.NewHashServer(conf, cacheStorage)
	pb.RegisterHashServiceServer(grpcServer, &hashServer)

	if err := grpcServer.Serve(lis); err != nil {
		lgr.ErrorWithExit(fmt.Sprintf("failed to serve: %s", err))
	}
}
