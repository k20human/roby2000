package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// Request logger infos
const (
	InfoPath      = "path"
	InfoStatus    = "status"
	InfoMethod    = "method"
	InfoQuery     = "query"
	InfoIP        = "ip"
	InfoUserAgent = "user_agent"
	InfoTime      = "time"
	InfoLatency   = "latency"
)

func RequestLogger(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		l.With(zap.Any(InfoPath, path)).
			With(zap.Any(InfoStatus, fmt.Sprintf("%d", c.Writer.Status()))).
			With(zap.Any(InfoMethod, c.Request.Method)).
			With(zap.Any(InfoQuery, query)).
			With(zap.Any(InfoIP, c.ClientIP())).
			With(zap.Any(InfoUserAgent, c.Request.UserAgent())).
			With(zap.Any(InfoTime, end.UTC().Format(time.RFC3339))).
			With(zap.Any(InfoLatency, latency.String())).
			Debug("")
	}
}
