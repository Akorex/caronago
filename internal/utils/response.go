package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Message any `json:"message"`
	Data    any    `json:"data"`
}

func SendSuccess(c *gin.Context, status int, message string, data any){
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data: data,
	})
}

func SendError(c *gin.Context, status int, message any){
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}