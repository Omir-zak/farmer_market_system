package routes

import (
	"database/sql"
	"farmer_market/controllers"
	"farmer_market/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserRoutes(router *gin.Engine, db *sql.DB) {
	// Страница с выбором роли
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// Форма регистрации покупателя
	router.GET("/register/buyer", func(c *gin.Context) {
		c.HTML(http.StatusOK, "buyer_form.html", nil)
	})

	// Форма регистрации фермера
	router.GET("/register/farmer", func(c *gin.Context) {
		c.HTML(http.StatusOK, "farmer_form.html", nil)
	})

	// Регистрация покуп	ателя
	router.POST("/register/buyer", func(c *gin.Context) {
		var buyer models.User
		if err := c.ShouldBind(&buyer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		err := controllers.RegisterBuyer(db, buyer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register buyer"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Buyer registered successfully"})
	})

	// Регистрация фермера
	router.POST("/register/farmer", func(c *gin.Context) {
		var farmer models.User

		// Попытка привязать JSON к модели
		if err := c.ShouldBindJSON(&farmer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid input: %s", err.Error()),
			})
			return
		}

		// Вызываем контроллер для регистрации фермера
		err := controllers.RegisterFarmer(db, farmer)
		if err != nil {
			if err.Error() == "email already exists" {
				c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register farmer"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Farmer registered successfully, waiting for admin approval"})
	})

}
