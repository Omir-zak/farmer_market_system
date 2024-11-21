package routes

/*import (
	"database/sql"
	"farmer_market/controllers"
	"farmer_market/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := controllers.CreateUser(db, user)
		if err != nil {
			if err.Error() == "email already exists" {
				// Correctly set the status to 409 Conflict
				c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to register user",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	})

	return router
}
*/
