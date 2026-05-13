package routes

import (
	"net/http"

	"github.com/Akorex/caronago/internal/handlers"
	"github.com/Akorex/caronago/internal/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handlers){
	v1 := r.Group("/api/v1")
	{
		RegisterAuthRoutes(v1, h.Auth)

		v1.GET("/health", func (c *gin.Context){
		utils.SendSuccess(c, http.StatusOK, "API is running successfully", nil)
	})
	}
}