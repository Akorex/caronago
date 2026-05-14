package middleware

import (
	"time"

	"github.com/Akorex/caronago/internal/utils"
	"github.com/gin-gonic/gin"
)

func StructuredLogger() gin.HandlerFunc{
	return func (c *gin.Context){
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		logger := utils.Log(c)
		latency := time.Since(start)
		status  := c.Writer.Status()

		if status >= 500{
			logger.Error("Server error", "status", status, "method", method, "path", path, "latency", latency.String(), "errors", c.Errors.String())
		}else if status >= 400{
			logger.Warn("Client error", "status", status, "method", method, "path", path, "latency", latency.String(), "errors", c.Errors.String())
		}else{
			logger.Info("Request completed", "status", status, "method", method, "path", path, "latency", latency.String())
		}
	}
}