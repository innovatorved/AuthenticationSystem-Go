package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/innovatorved/authenticationSystem-golang/controllers"
	"github.com/innovatorved/authenticationSystem-golang/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
