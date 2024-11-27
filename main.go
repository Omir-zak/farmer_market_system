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
	routes.RegisterFarmerRoutes(router, db)

	router.Run(":8080")
}
