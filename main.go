package main

import (
	routes "authenticationSystem-golang/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.Use(gin.Logger())

}
