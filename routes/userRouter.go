package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jerinthomas1404/go-jwt/controllers"
	"github.com/jerinthomas1404/go-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
