package main

import (
    "github.com/gin-gonic/gin"
	"DFHousing/DFHousingAPI/configs"
	"DFHousing/DFHousingAPI/routes"
	"net/http"
)



func main() {
    router := gin.Default()
    router.Use(corsMiddleware())

    configs.ConnectDB()

    // Define the `protected` group with JwtAuthMiddleware
    protected := router.Group("/api")

    // Use the protected group in other route files
    routes.UserRoute(protected)
    routes.PropertyRoute(protected)
    routes.ReviewRoute(protected)
    routes.RegisterRoute(router)
    routes.LoginRoute(router)

    router.Run("localhost:8080")
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with your frontend URL
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight requests
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }

        c.Next()
    }
}