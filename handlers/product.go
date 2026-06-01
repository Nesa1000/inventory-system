package handlers

import (
	"net/http"

	"github.com/Nesa1000/inventory-system/config"
	"github.com/Nesa1000/inventory-system/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Step 1: Read JSON body into product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// Step 2: Basic validation
	if product.Name == "" || product.SKU == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name and SKU are required",
		})
		return
	}

	// Step 3: Save to DB
	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	// Step 4: Return the created product
	c.JSON(http.StatusCreated, product)
}
