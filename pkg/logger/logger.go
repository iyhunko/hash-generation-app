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

type Log struct {
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
	l := Log{lgr: zapLogger}

	return &l, nil
}

func (l *Log) Error(messages ...string) {
	lgr := l.lgr
	message := strings.Join(messages, " ")

	lgr.Error(message)
}

func (l *Log) FatalError(err error) {
	l.Error(fmt.Sprintf("%v", err))
	os.Exit(1)
}

func (l *Log) Info(messages ...string) {
	lgr := l.lgr
	message := strings.Join(messages, " ")
	lgr.Info(message)
}

func (l *Log) Warn(messages ...string) {
	lgr := l.lgr
	message := strings.Join(messages, " ")
	lgr.Warn(message)
}

func (l *Log) WithStackTrace(errMsg string) Logger {
	return &Log{lgr: l.lgr.With(
		zap.String(stacktraceParam, errMsg),
	)}
}
