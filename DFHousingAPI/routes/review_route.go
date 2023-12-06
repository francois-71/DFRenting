package routes

import (
	"github.com/gin-gonic/gin"
	"DFHousing/DFHousingAPI/controllers"
	"DFHousing/DFHousingAPI/middlewares"
)

func ReviewRoute(router *gin.RouterGroup)  {
	router.POST("/createreview", middlewares.JwtAuthMiddleware(), controllers.CreateReview())
	router.GET("/review/:reviewId", middlewares.JwtAuthMiddleware(), controllers.GetAReview())
	router.PUT("/review/:reviewId", middlewares.JwtAuthMiddleware(), controllers.EditAReview()) 
	router.DELETE("/review/:reviewId", middlewares.JwtAuthMiddleware(), controllers.DeleteAReview())
	router.GET("/reviews/:propertyId", middlewares.JwtAuthMiddleware(), controllers.GetReviewsByPropertyId())
}