package middlewares

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"DFHousing/DFHousingAPI/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CheckAdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRole, tokenErr := token.ExtractTokenRole(c)
		fmt.Printf("tokenRole: %v\n", tokenRole)
		if tokenErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			c.Abort()
			return
		}

		//check if the role is admin
		
		if tokenRole != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}