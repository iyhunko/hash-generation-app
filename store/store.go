package store

import (
	"github.com/google/uuid"
	"os"
	"time"
)

type Storage interface {
	Get(key string) string
	Set(key string, v uuid.UUID) error
}

type Store struct {
	ttl time.Duration
}

func NewStore() Store {
	return Store{}
}

func (s *Store) Get(filePath string) []byte {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	return fileContent
}

func (s *Store) Set(filePath string, v []byte) error {
	err := os.WriteFile(filePath, v, 0644)
	if err != nil {
		return err
	}

	return nil
}
