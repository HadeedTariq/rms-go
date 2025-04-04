package main

import (
	"fmt"
	"log"
	"rms-platform/database"
	"rms-platform/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDb()
	database.MigrateTables()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Restaurant Management API!",
		})
	})
	api := r.Group("/api")
	routes.UserRoutes(api)
	routes.MenuRoutes(api)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting the Gin server: %v", err)
	}

	fmt.Println("Gin server is running on http://localhost:8080")
}
