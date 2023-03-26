package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	stacktraceParam    = "logger"
	initialSampling    = 100
	thereafterSampling = 100
)

type Logger interface {
	Warn(messages ...string)
	Error(messages ...string)
	FatalError(err error)
	Info(messages ...string)
	WithStackTrace(directory string) Logger
}

type logger struct {
	lgr *zap.Logger
}

func New() (Logger, error) {
	zapLogger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    initialSampling,
			Thereafter: thereafterSampling,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to init zap logger: %w", err)
	}
	l := &logger{lgr: zapLogger}

	return l, nil
}

func (l *logger) Error(messages ...string) {
	lgr := l.lgr
	message := strings.Join(messages, " ")

	lgr.Error(message)
}

func (l *logger) FatalError(err error) {
	l.Error(fmt.Sprintf("%v", err))
	os.Exit(1)
}

func (l *logger) Info(messages ...string) {
	lgr := l.lgr
	message := strings.Join(messages, " ")
	lgr.Info(message)
}

func (l *logger) Warn(messages ...string) {
	lgr := l.lgr
	message := strings.Join(messages, " ")
	lgr.Warn(message)
}

func (l *logger) WithStackTrace(errMsg string) Logger {
	return &logger{lgr: l.lgr.With(
		zap.String(stacktraceParam, errMsg),
	)}
}
