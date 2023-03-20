package store

import (
	"github.com/google/uuid"
	"github.com/iyhunko/hash-generation-app/logger"
	"os"
)

type Storage interface {
	Get(key string) string
	Set(key string, v uuid.UUID) error
}

type Store struct {
	log logger.Logger
}

func NewStore(log logger.Logger) Store {
	return Store{log: log}
}

func (s *Store) Get(filePath string) []byte {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		s.log.Warn(err.Error())
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
