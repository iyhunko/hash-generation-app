package http

import (
	config2 "github.com/iyhunko/hash-generation-app/internal/config"
	store2 "github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouter_Init(t *testing.T) {
	config := config2.NewConfig(nil)
	log, _ := logger.New()
	store := store2.NewStore(log)

	t.Run("new_router_created", func(t *testing.T) {
		router := InitRouter(config, store)
		assert.NotEmpty(t, router)
	})
}
