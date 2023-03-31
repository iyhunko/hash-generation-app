package main

import (
	"context"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/proto"
	"github.com/iyhunko/hash-generation-app/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	dLog "log"
)

const (
	startingClientMsg       = "Starting GRPC api client"
	failedToConnectErrMsg   = "failed to connect: %v"
	failedToFetchHashErrMsg = "failed to fetch hash: %v"
	resultMsg               = "Uuid: %s, Time: %s"
)

var lgr logger.Logger
var conf config.Config

func init() {
	nl, err := logger.New()
	if err != nil {
		dLog.Fatalf("failed to init create logger: %v.", err)
	}
	lgr = nl

	conf = config.NewConfig(lgr)
}

// This client is needed only for grpc_api testing
func main() {
	lgr.Info(startingClientMsg)
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
