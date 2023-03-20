package entity

import (
	"github.com/google/uuid"
	"time"
)

type Hash struct {
	Hash        uuid.UUID `json:"hash"`
	GeneratedAt time.Time `json:"generated_at"`
}

func NewHash() Hash {
	return Hash{
		GeneratedAt: time.Now(),
		Hash:        uuid.New(),
	}
}
