package main

import (
	"context"
	"fmt"
	"github.com/iyhunko/hash-generation-app/config"
	pb "github.com/iyhunko/hash-generation-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.Println("Starting GRPC api client")

	// load env variables to the Config struct
	var conf config.Config
	config.ReadEnv(&conf)

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		fmt.Sprintf(":%s", conf.GRPCServerPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHashServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), conf.ReadTimeout)
	defer cancel()
	r, err := c.GetHash(ctx, &pb.Hash{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Uuid: %s", r.GetUuid())
	log.Printf("Time: %s", r.GetTime())
}
