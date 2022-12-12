package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mayurkhairnar2525/restaurantManagement/middleware"
	"github.com/mayurkhairnar2525/restaurantManagement/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	// Router's
	routes.OrderRoutes(router)

	router.Run(":" + port)

}
