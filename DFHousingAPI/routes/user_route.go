package routes

import (
	"github.com/gin-gonic/gin"
	"DFHousing/DFHousingAPI/controllers"
    "DFHousing/DFHousingAPI/middlewares"
)

func UserRoute(router *gin.RouterGroup) {
    router.POST("/user",middlewares.JwtAuthMiddleware(), controllers.CreateUser())
    router.GET("/user/:userId", middlewares.JwtAuthMiddleware(), controllers.GetAUser())
    router.GET("/user/isadmin", middlewares.JwtAuthMiddleware(), controllers.IsUserAdmin())
    router.PUT("/user/:userId", middlewares.JwtAuthMiddleware(), controllers.EditAUser())
    router.DELETE("/user/:userId", middlewares.JwtAuthMiddleware(),controllers.DeleteAUser())
    router.GET("/users", middlewares.JwtAuthMiddleware(), controllers.GetAllUsers())
    router.GET("/user", middlewares.JwtAuthMiddleware(), controllers.GetCurrentUser())
}