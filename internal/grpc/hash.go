package grpc

import (
	"context"
	"encoding/json"
	"fmt"
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
func (s *HashServer) GetHash(ctx context.Context, pHash *hash.Hash) (*hash.Hash, error) {
	hashBytes := s.store.Get(s.config.HashFilePath)
	if hashBytes == nil {
		return pHash, nil
	}
	fHash := entity.Hash{}
	err := json.Unmarshal(hashBytes, &fHash)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file bytes: %w", err)
	}
	pHash.Uuid = fHash.Hash.String()
	pHash.Time = fHash.GeneratedAt.String()

	return pHash, nil
}
