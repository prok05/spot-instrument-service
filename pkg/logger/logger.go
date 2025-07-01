package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Interface interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type Logger struct {
	logger *zap.SugaredLogger
}

func New(level string) (*Logger, error) {
	cfg := zap.Config{
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	switch strings.ToLower(level) {
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "warn":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("logger - NewLogger - zap.cfg.Build: %w", err)
	}
	sugar := logger.Sugar()

	return &Logger{
		logger: sugar,
	}, nil
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debugw(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Infow(msg, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warnw(message, args...)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.logger.Errorw(message, args...)
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	l.logger.Fatalw(message, args...)
}
