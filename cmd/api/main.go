package main

import (
	"log"
	"net/http"

	"github.com/Akorex/caronago/internal/config"
	"github.com/Akorex/caronago/internal/database"
	"github.com/Akorex/caronago/internal/handlers"
	"github.com/Akorex/caronago/internal/routes"
	"github.com/Akorex/caronago/internal/services"
	"github.com/Akorex/caronago/internal/utils"
	"github.com/gin-gonic/gin"
)

func main(){
	cfg := config.LoadConfig()

	r := gin.New()
	r.Use(gin.Logger())

	database.Connect(cfg.DBUrl)

	authService := services.NewAuthService(database.DB)
	authHandler := handlers.NewAuthHandler(authService)

	v1 := r.Group("/api/v1")
	routes.RegisterAuthRoutes(v1, authHandler)

	r.GET("/api/v1/health", func (c *gin.Context){
		utils.SendSuccess(c, http.StatusOK, "API is running successfully", nil)
	})

	port := cfg.Port
	if port == ""{
		port = "8080"
	}

	r.Run(":" + port)



	log.Println("Server running on port", port)


}