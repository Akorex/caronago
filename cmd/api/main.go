package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.New()
	r.Use(gin.Logger())


	r.Run(":8080")

	log.Println("Server running on port", 8080)


}