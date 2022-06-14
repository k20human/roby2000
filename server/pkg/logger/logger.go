package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

func New() (*zap.Logger, error) {
	var err error
	var l *zap.Logger
	var c *config
	var level zapcore.Level

	if c, err = initConfig(); err != nil {
		log.Fatalf("Configuration: %s\n", err)
	}

	if level, err = zapcore.ParseLevel(c.Level); err != nil {
		log.Fatalf(err.Error())
	}

	if l, err = (zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         "console",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()); err != nil {
		return nil, err
	}

	return l, nil
}

func Close(l *zap.Logger) {
	if err := l.Sync(); err != nil {
		if strings.Contains(err.Error(), "bad file descriptor") || strings.Contains(err.Error(), "invalid argument") {
			// ignore because the stderr should not sync.
			return
		}
		l.Error(err.Error())
	}
}
