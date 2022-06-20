package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jerinthomas1404/go-jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
