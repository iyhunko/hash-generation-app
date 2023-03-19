package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/iyhunko/hash-generation-app/config"
	pb "github.com/iyhunko/hash-generation-app/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHashServiceServer
}

// GetHash returns the current hash
func (s *server) GetHash(ctx context.Context, hash *pb.Hash) (*pb.Hash, error) {
	id := uuid.New()
	tm := time.Now().Format(time.RFC1123)

	return &pb.Hash{Time: tm, Uuid: id.String()}, nil
}

func main() {
	log.Println("Starting GRPC api server")

	// load env variables to the Config struct
	var conf config.Config
	config.ReadEnv(&conf)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHashServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
