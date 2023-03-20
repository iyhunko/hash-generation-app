package grpc

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/iyhunko/hash-generation-app/proto"
	"time"
)

type Server struct {
	pb.UnimplementedHashServiceServer
}

// GetHash returns the current hash
func (s *Server) GetHash(ctx context.Context, hash *pb.Hash) (*pb.Hash, error) {
	id := uuid.New()
	tm := time.Now().Format(time.RFC1123)

	return &pb.Hash{Time: tm, Uuid: id.String()}, nil
}
