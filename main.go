package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/track-the-trails/config"
	"github.com/yourusername/track-the-trails/routes"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
