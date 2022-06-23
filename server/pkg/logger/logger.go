package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	logPath    = "/var/log/roby2000/roby2000.log"
	fodlerPerm = 0o777
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

	if err = os.MkdirAll(filepath.Dir(logPath), fodlerPerm); err != nil {
		return nil, errors.Errorf("unable to create directory for file '%s': %s\nPlease create it manually\n", logPath, err.Error())
	}

	if l, err = (zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         "console",
		OutputPaths:      []string{"stdout", logPath},
		ErrorOutputPaths: []string{"stderr", logPath},
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
