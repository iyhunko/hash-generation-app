package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_NewConfig(t *testing.T) {
	t.Run("default_values", func(t *testing.T) {
		conf := NewConfig(nil)
		assert.Equal(t, "8001", conf.HTTPServerPort)
		assert.Equal(t, "8002", conf.GRPCServerPort)
		assert.Equal(t, "15s", conf.ReadTimeout.String())
		assert.Equal(t, "15s", conf.WriteTimeout.String())
		assert.Equal(t, "hash.json", conf.HashFilePath)
		assert.Equal(t, "5s", conf.HashGenerationInterval.String())
	})
}
