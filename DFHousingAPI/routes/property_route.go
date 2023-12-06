package routes

import (
	"github.com/gin-gonic/gin"
	"DFHousing/DFHousingAPI/controllers"
	"DFHousing/DFHousingAPI/middlewares"
)

func PropertyRoute(router *gin.RouterGroup)  {
	
	router.POST("/createproperty", middlewares.JwtAuthMiddleware(), controllers.CreateProperty())
	router.GET("/property/:propertyId", middlewares.JwtAuthMiddleware(), controllers.GetAProperty())
	router.GET("/properties", controllers.GetAllProperties())
	router.PATCH("properties/approve/:propertyId", middlewares.JwtAuthMiddleware(), middlewares.CheckAdminRoleMiddleware(), controllers.SetPropertyTrueApproval())
	router.PATCH("properties/reject/:propertyId", middlewares.JwtAuthMiddleware(), middlewares.CheckAdminRoleMiddleware(), controllers.SetPropertyFalseApproval())
	router.GET("/properties/requireapproval", middlewares.JwtAuthMiddleware(), middlewares.CheckAdminRoleMiddleware(), controllers.GetPropertiesByFalseApproval())
	/*
	router.PUT("/property/:propertyId", controllers.EditAProperty()) 
	router.DELETE("/property/:propertyId", controllers.DeleteAProperty())
	
	router.GET("/properties/:userId", controllers.GetPropertiesByHost())
	router.GET("/property/:propertyId/reviews", controllers.GetReviewsByPropertyId())
	*/
}