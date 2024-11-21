package routes

import (
	"database/sql"
	"farmer_market/controllers"
	"farmer_market/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterAdminRoutes(router *gin.Engine, db *sql.DB) {
	// Вход администратора
	router.GET("/admin/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_login.html", nil)
	})

	router.POST("/admin/login", func(c *gin.Context) {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Проверка тела запроса
		if err := c.ShouldBind(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Вызов функции LoginAdmin
		token, err := controllers.LoginAdmin(db, credentials.Email, credentials.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Возврат JWT токена при успешном входе
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
		})
	})

	// Панель управления администратора
	/*router.GET("/admin/dashboard", middlewares.AdminAuthMiddleware(), func(c *gin.Context) {

	})*/

	// Получение заявок фермеров
	router.GET("/admin/pending", middlewares.AdminAuthMiddleware(), func(c *gin.Context) {
		farmers, err := controllers.GetPendingFarmers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch pending farmers: %v", err.Error())})
			return
		}
		c.JSON(http.StatusOK, farmers)
	})

	// Подтверждение фермера
	router.POST("/admin/farmers/approve/:id", middlewares.AdminAuthMiddleware(), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid farmer ID"})
			return
		}
		err = controllers.ApproveFarmer(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve farmer"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Farmer approved successfully"})
	})

	// Отклонение фермера
	router.POST("/admin/farmers/reject/:id", middlewares.AdminAuthMiddleware(), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid farmer ID"})
			return
		}
		err = controllers.RejectFarmer(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject farmer"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Farmer rejected successfully"})
	})
}
