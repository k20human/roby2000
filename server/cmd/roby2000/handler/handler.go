package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/k20human/roby2000/pkg/logger"
	"github.com/k20human/roby2000/pkg/robot"
	"go.uber.org/zap"
)

type Handler interface {
	Move(c *gin.Context)
	Close()
}

type handler struct {
	logger *zap.Logger
	driver robot.Robot
}

func New() (*handler, error) {
	var err error
	var h handler

	if h.logger, err = logger.New(); err != nil {
		return nil, err
	}

	if h.driver, err = robot.New(); err != nil {
		return nil, err
	}

	return &h, nil
}

// Close drivers.
func (h *handler) Close() {
	if err := h.driver.Close(); err != nil {
		h.logger.Error("Error during stop robot process", zap.Error(err))
	}

	logger.Close(h.logger)
}
