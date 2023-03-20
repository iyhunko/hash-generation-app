package grpc

import (
	config2 "github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/proto"
	store2 "github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash_GetHash(t *testing.T) {
	config := config2.NewConfig(nil)
	log, _ := logger.New()
	store := store2.NewStore(log)

	t.Run("returns_empty_hash_if_no_file_present", func(t *testing.T) {
		protoHash := hash.Hash{}
		hashServer := NewHashServer(config, store)
		_, err := hashServer.GetHash(nil, &protoHash)
		assert.Empty(t, err)
		assert.Empty(t, protoHash.Time)
		assert.Empty(t, protoHash.Uuid)
	})
}
