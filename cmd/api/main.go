package main

import (
	"log"

	"github.com/Akorex/caronago/internal/config"
	"github.com/Akorex/caronago/internal/database"
	"github.com/Akorex/caronago/internal/handlers"
	"github.com/Akorex/caronago/internal/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	cfg := config.LoadConfig()

	r := gin.New()
	r.Use(gin.Logger())

	database.Connect(cfg.DBUrl)
	
	h := handlers.RegisterHandlers(database.DB, cfg)
	routes.RegisterRoutes(r, h)

	port := cfg.Port
	if port == ""{
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}