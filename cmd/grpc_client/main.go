package main

import (
	"context"
	"fmt"
	"github.com/iyhunko/hash-generation-app/config"
	"github.com/iyhunko/hash-generation-app/logger"
	pb "github.com/iyhunko/hash-generation-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// This client is needed only for grpc_api testing
func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	lgr.Info("Starting GRPC api client")
	conf := config.NewConfig(lgr)

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		fmt.Sprintf(":%s", conf.GRPCServerPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		lgr.ErrorWithExit(fmt.Sprintf("Failed to connect: %v", err))
	}
	defer conn.Close()

	c := pb.NewHashServiceClient(conn)
	// Contact the server and log out its response.
	ctx, cancel := context.WithTimeout(context.Background(), conf.ReadTimeout)
	defer cancel()
	r, err := c.GetHash(ctx, &pb.Hash{})
	if err != nil {
		lgr.ErrorWithExit(fmt.Sprintf("Failed to fetch hash: %v", err))
	}

	lgr.Info(fmt.Sprintf("Uuid: %s, Time: %s", r.GetUuid(), r.GetTime()))
}
