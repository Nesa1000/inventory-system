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

func GetProducts(c *gin.Context) {
	var products []models.Product

	// Fetch all products
	result := config.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch products",
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Check product exists
	result := config.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Check product exists
	result := config.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Read input as a map so only provided fields are updated
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// Update only the fields provided
	result = config.DB.Model(&product).Updates(input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update product",
		})
		return
	}

	// Return the updated product
	config.DB.First(&product, id)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Check product exists
	result := config.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Delete the product
	result = config.DB.Unscoped().Delete(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
