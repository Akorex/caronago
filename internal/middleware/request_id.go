package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc{
	return func (c *gin.Context){
		id := c.GetHeader("X-Request-ID")

		if id == ""{
			id =uuid.New().String()
		}

		c.Set("requestId", id)
		c.Header("X-Request-ID", id)

		c.Next()
	}
}