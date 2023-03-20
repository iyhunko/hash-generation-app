package grpc

import (
	"context"
	"encoding/json"
	"github.com/iyhunko/hash-generation-app/config"
	"github.com/iyhunko/hash-generation-app/entity"
	pb "github.com/iyhunko/hash-generation-app/proto"
	"github.com/iyhunko/hash-generation-app/store"
)

type HashServer struct {
	pb.UnimplementedHashServiceServer
	config config.Config
	store  store.Store
}

func NewHashServer(
	config config.Config,
	store store.Store,
) HashServer {
	return HashServer{
		config: config,
		store:  store,
	}
}

// GetHash returns the current hash
func (s *HashServer) GetHash(ctx context.Context, iHash *pb.Hash) (*pb.Hash, error) {
	hashBytes := s.store.Get(s.config.HashFilePath)
	fHash := entity.Hash{}
	err := json.Unmarshal(hashBytes, &fHash)
	if err != nil {
		return nil, err
	}
	iHash.Uuid = fHash.Hash.String()
	iHash.Time = fHash.GeneratedAt.String()

	return iHash, nil
}
