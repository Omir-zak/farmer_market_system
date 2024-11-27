package routes

import (
	"database/sql"
	"farmer_market/controllers"
	"farmer_market/middlewares"
	"farmer_market/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterFarmerRoutes(router *gin.Engine, db *sql.DB) {
	// Группа маршрутов для фермеров с аутентификацией
	farmerGroup := router.Group("/farmer", middlewares.FarmerAuthMiddleware())

	// Получение всех продуктов фермера
	farmerGroup.GET("/products", func(c *gin.Context) {
		farmerID := c.GetInt("user_id")
		products, err := controllers.GetFarmerProducts(db, farmerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	// Добавление нового продукта
	farmerGroup.POST("/products", func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		product.FarmerID = c.GetInt("user_id")

		if err := controllers.AddProduct(db, &product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
	})

	// Обновление продукта
	farmerGroup.PUT("/products/:id", func(c *gin.Context) {
		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		product.FarmerID = c.GetInt("user_id")

		if err := controllers.UpdateProduct(db, productID, &product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
	})

	// Удаление продукта
	farmerGroup.DELETE("/products/:id", func(c *gin.Context) {
		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		if err := controllers.DeleteProduct(db, productID, c.GetInt("user_id")); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	})
}
