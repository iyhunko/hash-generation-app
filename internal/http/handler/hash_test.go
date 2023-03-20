package handler

import (
	config2 "github.com/iyhunko/hash-generation-app/internal/config"
	store2 "github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHashHandler_Get(t *testing.T) {
	config := config2.NewConfig(nil)
	log, _ := logger.New()
	store := store2.NewStore(log)
	hash := NewHashHandler(config, store)

	req, _ := http.NewRequest(http.MethodGet, "/hash", nil)
	rr := httptest.NewRecorder()

	t.Run("response_is_server_error", func(t *testing.T) {
		hash.Get(rr, req)
		assert.Equal(t, 500, rr.Code)
		assert.NotEmpty(t, failedToFetchMsg, rr.Body.String())
	})
}
