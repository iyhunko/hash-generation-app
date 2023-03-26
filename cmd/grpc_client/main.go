package main

import (
	"context"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/proto"
	"github.com/iyhunko/hash-generation-app/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	startingClientMsg       = "Starting GRPC api client"
	failedToConnectErrMsg   = "failed to connect: %v"
	failedToFetchHashErrMsg = "failed to fetch hash: %v"
	resultMsg               = "Uuid: %s, Time: %s"
)

// This client is needed only for grpc_api testing
func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatalf("Failed to init create logger: %v.", err)
	}

	lgr.Info(startingClientMsg)
	conf := config.NewConfig(lgr)

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		fmt.Sprintf(":%s", conf.GRPCServerPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		lgr.FatalError(fmt.Errorf(failedToConnectErrMsg, err))
	}
	defer conn.Close()

	c := hash.NewHashServiceClient(conn)
	// Contact the server and log out its response.
	ctx, cancel := context.WithTimeout(context.Background(), conf.ReadTimeout)
	defer cancel()
	r, err := c.GetHash(ctx, &hash.Hash{})
	if err != nil {
		lgr.FatalError(fmt.Errorf(failedToFetchHashErrMsg, err))
	}

	lgr.Info(fmt.Sprintf(resultMsg, r.GetUuid(), r.GetTime()))
}
