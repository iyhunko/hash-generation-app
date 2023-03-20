package store

import (
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/bluele/gcache"
)

type Cache interface {
	Get(key string) string
	Set(key string, v uuid.UUID) error
}

type Store struct {
	cache gcache.Cache
	ttl   time.Duration
}

func NewStore(size int, ttl time.Duration) Store {
	return Store{
		cache: gcache.New(size).Simple().Build(),
		ttl:   ttl,
	}
}

func (s *Store) Get(key string) []byte {
	v, err := s.cache.Get(key)
	if err != nil {
		fmt.Println("Failed to load from cache:", err)
		return nil
	}
	return v.([]byte)
}

func (s *Store) Set(key string, v []byte) error {
	err := s.cache.SetWithExpire(key, v, s.ttl)
	if err != nil {
		return err
	}
	return nil
}
