package store

import (
	"github.com/golang/mock/gomock"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"github.com/iyhunko/hash-generation-app/pkg/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore_Get(t *testing.T) {
	t.Run("empty_result", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockLogger := mock.NewMockLogger(mockCtrl)
		mockLogger.EXPECT().Warn("open test_file_path.json: no such file or directory")

		storeWithMock := NewStore(mockLogger)
		r := storeWithMock.Get("test_file_path.json")

		assert.Nil(t, r)
	})

	log, _ := logger.New()
	store := NewStore(log)
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
