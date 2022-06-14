package handler

import (
	"github.com/k20human/roby2000/pkg/logger"
	"go.uber.org/zap"
)

type Handler interface {
	Close()
}

type handler struct {
	logger *zap.Logger
}

func New() (*handler, error) {
	var err error
	var h handler

	if h.logger, err = logger.New(); err != nil {
		return nil, err
	}

	return &h, nil
}

// Close drivers.
func (h *handler) Close() {
	logger.Close(h.logger)
}
