package routes

import (
	"github.com/Akorex/caronago/internal/config"
	"github.com/Akorex/caronago/internal/handlers"
	"github.com/Akorex/caronago/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, h *handlers.AuthHandler, cfg *config.Config){
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
	}


	protected := router.Group("/auth")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.GET("/me", h.GetMe)
	}
}