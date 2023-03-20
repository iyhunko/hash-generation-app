package logger

import (
	"testing"
)

var (
	loggerErrorStr = "some error"
)

func TestLogger_Error(t *testing.T) {
	lgr, _ := New()

	t.Run("success", func(t *testing.T) {
		lgr.Error(loggerErrorStr)
	})
}

func TestLogger_Info(t *testing.T) {
	lgr, _ := New()

	t.Run("success", func(t *testing.T) {
		lgr.Info(loggerErrorStr)
	})
}

func TestLogger_Warn(t *testing.T) {
	lgr, _ := New()

	t.Run("success", func(t *testing.T) {
		lgr.Warn(loggerErrorStr)
	})
}
