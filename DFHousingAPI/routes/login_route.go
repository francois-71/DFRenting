package routes

import (
	"github.com/gin-gonic/gin"
	"DFHousing/DFHousingAPI/controllers"
)

func LoginRoute(router *gin.Engine)  {
	
	router.POST("/login", controllers.LoginUser())
	router.GET("/isloggedin", controllers.IsLoggedIn())
}