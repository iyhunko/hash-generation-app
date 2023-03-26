package main

import (
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain_RefreshHash(t *testing.T) {
	lgr, _ := logger.New()
	conf := config.NewConfig(lgr)
	storage := store.NewStore(lgr)
	conf.HashFilePath = "test_hash.json"

	t.Run("success_hash_refresh", func(t *testing.T) {
		err := refreshHash(lgr, conf, storage)
		assert.Empty(t, err)

		_, err = os.OpenFile(conf.HashFilePath, os.O_RDWR, 0644)
		assert.Nil(t, err)
	})

	t.Run("fail_empty_path", func(t *testing.T) {
		conf.HashFilePath = ""

		err := refreshHash(lgr, conf, storage)
		assert.Equal(t, "failed to set value to storage: failed to set value to store: open : no such file or directory", err.Error())
	})
}
