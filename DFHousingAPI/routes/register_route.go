package routes

import (
	"github.com/gin-gonic/gin"
	"DFHousing/DFHousingAPI/controllers"
)

func RegisterRoute(router *gin.Engine)  {
	router.POST("/register", controllers.RegisterUser())
}