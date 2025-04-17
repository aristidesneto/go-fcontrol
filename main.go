package main

import (
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	configs.ConnectDB()

	routes := gin.Default()
	routers.InitRouter(routes)

	log.Println("Server is running on port 8080")
	if err := routes.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
