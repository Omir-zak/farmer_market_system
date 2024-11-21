package main

import (
	"farmer_market/config"
	"farmer_market/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*") // Загрузка HTML-шаблонов

	routes.RegisterUserRoutes(router, db)
	routes.RegisterAdminRoutes(router, db)

	router.Run(":8080")
}

/*package main

import (
	"farmer_market/config"
	"farmer_market/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()
	defer config.DB.Close()

	// Initialize Gin routes
	router := gin.Default()

	//Register user routes
	routes.RegisterUserRoutes(router, config.DB)

	// Start the server
	log.Println("Starting server on :8080...")
	router.Run(":8080")
}
*/
