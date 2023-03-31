package store

import (
	"fmt"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"os"
)

type Storage interface {
	Get(filePath string) []byte
	Set(filePath string, v []byte) error
}

type Store struct {
	log logger.Logger
}

func NewStore(log logger.Logger) Storage {
	return Store{log: log}
}

func (s Store) Get(filePath string) []byte {
	//TODO: should be refactored
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		s.log.Warn(err.Error())
		return nil
	}

	return fileContent
}

func (s Store) Set(filePath string, v []byte) error {
	err := os.WriteFile(filePath, v, 0644)
	if err != nil {
		return fmt.Errorf("failed to set value to store: %w", err)
	}

	return nil
}
