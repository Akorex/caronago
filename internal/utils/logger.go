package utils

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

var baseLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Log(c *gin.Context) *slog.Logger{
	if id, exists := c.Get("requestId"); exists{
		return baseLogger.With("requestId", id)
	}

	return baseLogger
}