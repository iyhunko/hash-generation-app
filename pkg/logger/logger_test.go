package logger

import (
	"testing"
)

var (
	loggerErrorStr = "some error"
)

func TestLogger_Error(t *testing.T) {
	lgr, _ := New()

	t.Run("log_error", func(t *testing.T) {
		lgr.Error(loggerErrorStr)
	})
}

func TestLogger_Info(t *testing.T) {
	lgr, _ := New()

	t.Run("log_info", func(t *testing.T) {
		lgr.Info(loggerErrorStr)
	})
}

func TestLogger_Warn(t *testing.T) {
	lgr, _ := New()

	t.Run("log_warn", func(t *testing.T) {
		lgr.Warn(loggerErrorStr)
	})
}
