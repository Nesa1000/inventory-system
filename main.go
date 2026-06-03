package main

import (
	"os"
	"github.com/Nesa1000/inventory-system/config"
	"github.com/Nesa1000/inventory-system/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env locally, ignore error in production
	godotenv.Load()

	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
