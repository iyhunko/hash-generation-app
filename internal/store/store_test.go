package store

import (
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore_Get(t *testing.T) {
	log, _ := logger.New()
	store := NewStore(log)

	t.Run("empty_result", func(t *testing.T) {
		r := store.Get("test_file_path.json")

		assert.Nil(t, r)
	})

	t.Run("empty_result", func(t *testing.T) {
		r := store.Get("test_file_path.json")

		assert.Nil(t, r)
	})

	t.Run("non_empty_result", func(t *testing.T) {
		r := store.Get("test_file.json")

		assert.Equal(t, "{\"hash\":\"de1b6672-6b12-4b8a-92a2-5cb0af28922f\",\"generated_at\":\"2023-03-20T16:00:29.311478+02:00\"}\n", string(r))
	})
}

func TestStore_Set(t *testing.T) {
	log, _ := logger.New()
	store := NewStore(log)

	t.Run("success_set", func(t *testing.T) {
		err := store.Set("test_file_path.json", []byte("some text data"))

		assert.Nil(t, err)
	})
}
