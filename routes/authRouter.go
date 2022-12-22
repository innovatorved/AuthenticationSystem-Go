package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/innovatorved/AuthenticationSystem-Go/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/register", controller.Register())
	incomingRoutes.POST("users/login", controller.Login())
}
