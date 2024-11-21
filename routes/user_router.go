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

	// Регистрация покупателя
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
		rawData, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		var farmer models.User
		if err := c.ShouldBind(&farmer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err), "req": string(rawData)})
			return
		}

		err = controllers.RegisterFarmer(db, farmer)
		if err != nil {
			if err.Error() == "email already exists" {
				c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register farmer", "details": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Farmer registered successfully, waiting for admin approval"})
	})
}

/*package routes

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
}*/
