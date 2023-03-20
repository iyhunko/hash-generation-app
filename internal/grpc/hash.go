package grpc

import (
	"context"
	"encoding/json"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/entity"
	"github.com/iyhunko/hash-generation-app/internal/proto"
	"github.com/iyhunko/hash-generation-app/internal/store"
)

type HashServer struct {
	hash.UnimplementedHashServiceServer
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
func (s *HashServer) GetHash(ctx context.Context, iHash *hash.Hash) (*hash.Hash, error) {
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
