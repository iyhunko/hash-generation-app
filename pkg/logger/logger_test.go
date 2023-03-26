package logger

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

var (
	loggerError = errors.New("some error")
)

func TestLogger_FatalError(t *testing.T) {
	lgr, _ := New()
	if os.Getenv("CRASH") == "1" {
		lgr.FatalError(loggerError)
		return
	}
	cmd := exec.Command(os.Args[0], "test -test.run=TestLogger_FatalError")
	cmd.Env = append(os.Environ(), "CRASH=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, expected exit status 1", err)
}

func TestLogger_Error(t *testing.T) {
	lgr, _ := New()

	t.Run("log_error", func(t *testing.T) {
		lgr.Error(loggerError.Error())
	})
}

func TestLogger_Info(t *testing.T) {
	lgr, _ := New()

	t.Run("log_info", func(t *testing.T) {
		lgr.Info(loggerError.Error())
	})
}

func TestLogger_Warn(t *testing.T) {
	lgr, _ := New()

	t.Run("log_warn", func(t *testing.T) {
		lgr.Warn(loggerError.Error())
	})
}
