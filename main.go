package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/innovatorved/authenticationSystem-golang/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for router"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for router api2"})
	})

	router.Run(":" + port)
}
