package middlewares

import (
	"farmer_market/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := &controllers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(controllers.JwtKey), nil
		})

		if err != nil || !token.Valid || claims.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func FarmerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Удаляем префикс "Bearer " из токена
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := &controllers.Claims{}

		// Парсим токен и проверяем его валидность
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(controllers.JwtKey), nil
		})
		if err != nil || !token.Valid || claims.Role != "farmer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Сохраняем ID фермера в контексте для последующего использования
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
