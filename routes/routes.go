package routes

import (
	"github.com/Nesa1000/inventory-system/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/products", handlers.CreateProduct)
		api.GET("/products", handlers.GetProducts)
		api.GET("/products/:id", handlers.GetProduct)
		api.PUT("/products/:id", handlers.UpdateProduct)
		api.DELETE("/products/:id", handlers.DeleteProduct)
	}
}
