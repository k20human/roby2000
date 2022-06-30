package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/k20human/roby2000/cmd/roby2000/handler"
	"github.com/k20human/roby2000/pkg/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	var l *zap.Logger
	var c *config

	if c, err = initConfig(); err != nil {
		log.Fatalf("Configuration: %s\n", err)
	}

	if l, err = logger.New(); err != nil {
		log.Fatalf("Logger: %s\n", err)
	}

	defer logger.Close(l)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	gin.SetMode(c.GinMode)
	engine, h := initRouter(l)

	defer h.Close()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal(err.Error())
		}
	}()

	l.Info("Server and robot driver started", zap.Any("port", c.Port))

	<-ctx.Done()

	stop()
	l.Info("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown", zap.Error(err))
	}

	l.Info("Server exiting")
}

func initRouter(l *zap.Logger) (*gin.Engine, handler.Handler) {
	var err error
	var h handler.Handler

	engine := gin.New()
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Use(gin.Recovery())

	if gin.Mode() != gin.ReleaseMode {
		engine.Use(handler.RequestLogger(l))
		l.Warn("Debug mode activated on Gin, log requests activated")
	}

	if h, err = handler.New(); err != nil {
		l.Fatal(err.Error())
	}

	initRoutes(engine.RouterGroup, h)

	return engine, h
}

func initRoutes(r gin.RouterGroup, h handler.Handler) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Roby2000")
	})

	r.GET("/move/:action", h.Move)
	r.GET("/light/:action/:type/:color", h.Light)
}
