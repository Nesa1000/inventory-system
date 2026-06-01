package main

import (
	"github.com/Nesa1000/inventory-system/config"
	"github.com/Nesa1000/inventory-system/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
