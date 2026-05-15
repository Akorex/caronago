package middleware

import (
	"time"

	"github.com/Akorex/caronago/internal/enums"
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
		latency := time.Since(start).Milliseconds()
		status  := c.Writer.Status()

		if status >= 500{
			logger.Error(enums.MsgServerError, "status", status, "method", method, "path", path, "latency", latency, "errors", c.Errors.String())
		}else if status >= 400{
			logger.Warn(enums.MsgClientError, "status", status, "method", method, "path", path, "latency", latency, "errors", c.Errors.String())
		}else{
			logger.Info(enums.MsgRequestCompleted, "status", status, "method", method, "path", path, "latency", latency)
		}
	}
}