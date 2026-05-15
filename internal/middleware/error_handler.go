package middleware

import (
	"net/http"

	"github.com/Akorex/caronago/internal/enums"
	"github.com/Akorex/caronago/internal/errors"
	"github.com/Akorex/caronago/internal/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next() // run the handler first

		if len(c.Errors) == 0{
			return
		}

		ginErr := c.Errors.Last()
		logger := utils.Log(c)

		if apiError, ok := ginErr.Err.(*errors.ApiError); ok{
			if apiError.Raw != nil{
				logger.Error(enums.MsgInternalError, "cause", apiError.Raw.Error())
			}

			utils.SendError(c, apiError.Code, apiError.Message)
			return
		}

		logger.Error(enums.MsgUnhandledError,"cause", ginErr.Error())
		utils.SendError(c, http.StatusInternalServerError, "Internal Server Error")
	}
	
}