package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akorex/caronago/internal/config"
	"github.com/Akorex/caronago/internal/database"
	"github.com/Akorex/caronago/internal/handlers"
	"github.com/Akorex/caronago/internal/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	cfg := config.LoadConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	database.Connect(cfg.DBUrl)
	
	h := handlers.RegisterHandlers(database.DB, cfg)
	routes.RegisterRoutes(r, h, cfg)

	port := cfg.Port
	if port == ""{
		port = "8080"
	}

	srv := &http.Server{
		Addr: ":" + port,
		Handler: r,
		ReadTimeout:  5 * time.Second,  // Max time to read the incoming request
		WriteTimeout: 10 * time.Second, // Max time to write the outgoing response
		IdleTimeout:  120 * time.Second, // Max time a connection can stay open without doing anything
	}


	// start the server in a goroutine 
	go func(){
		log.Printf("Server running on port %s", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Fatalf("Server failed to start :%v\n", err)
		}
	}()


	// wait for an interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Println("Shutdown signal received. Draining connections...")

	// create a 5-second timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil{
		log.Fatal("Server forced to shutdown forcefully: ", err)
	}

	log.Println("Server exited cleanly.")

}