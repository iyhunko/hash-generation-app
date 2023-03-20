package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash_NewHash(t *testing.T) {
	t.Run("new_hash_has_non_nil_values", func(t *testing.T) {
		hash := NewHash()
		assert.NotEmpty(t, hash.Hash)
		assert.NotEmpty(t, hash.GeneratedAt)
	})
}
