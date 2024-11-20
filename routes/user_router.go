package routes

import (
	"database/sql"
	"farmer_market/controllers"
	"farmer_market/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserRoutes(router *gin.Engine, db *sql.DB) {
	// User registration endpoint
	router.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid input",
				"details": err.Error(),
			})
			return
		}

		err := controllers.CreateUser(db, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to register user",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	})

	// User login endpoint
	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		user, err := controllers.GetUserByEmail(db, loginData.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Placeholder password check (add bcrypt validation later)
		if loginData.Password != user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
	})
}
