package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAndValidate[T any](c *gin.Context, input *T) bool{
	if err := c.ShouldBindJSON(input); err != nil{
		if err == io.EOF{
			SendError(c, http.StatusBadRequest, "Request body should not be empty")
			return false
		}


		if msg := HandleValidationError(err); msg != nil{
			SendError(c, http.StatusBadRequest, msg)
			return false
		}

		if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok{
			msg := fmt.Sprintf("Field %s must be of type '%s'", unmarshalErr.Field, unmarshalErr.Type.String())
			SendError(c,http.StatusBadRequest, msg)
			return false
		}

		log.Printf("[BindAndValidate Error] Path: %s | Err: %v", c.Request.URL.Path, err)
		SendError(c, http.StatusBadRequest, "Invalid request payload")
		return false
	}

	return true
}