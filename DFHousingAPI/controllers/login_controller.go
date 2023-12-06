package controllers

import (
	"context"
	"DFHousing/DFHousingAPI/models"
	"DFHousing/DFHousingAPI/responses"
	"DFHousing/DFHousingAPI/utils/token"
	validator "DFHousing/DFHousingAPI/controllers/validators"
	
	"net/http"
	"time"
	"fmt"
	

	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"os"
	// add jwt token auth
)

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var loginInfo models.LoginInput

		// Validate and extract email and password from request body
		if err := c.BindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// use validator library to validate required fields
		if validationErr := validator.Validate.Struct(loginInfo); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": "Invalid credentials"},
			})
			return
		}

		// Retrieve user data by email
		var user models.User
		filter := bson.M{"email": loginInfo.Email}
		err := userCollection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.UserResponse{
				Status:  http.StatusUnauthorized,
				Message: "error",
				Data:    map[string]interface{}{"data": "Invalid credentials"},
			})
			return
		}

		// Check if password matches
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.UserResponse{
				Status:  http.StatusUnauthorized,
				Message: "error",
				Data:    map[string]interface{}{"data": "Invalid credentials"},
			})
			return
		}

		// Generate JWT token for the user ID
		token, err := token.GenerateToken(user.Id.Hex(), user.Role)
		fmt.Printf("token: %v", token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Please try again later",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Respond with success and token
		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"token": token},
		})
	}
}

func IsLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := token.ExtractToken(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, responses.UserResponse{
				Status:  http.StatusUnauthorized,
				Message: "error",
				Data:    map[string]interface{}{"data": "Unauthorized"},
			})
			return
		}

		// Parse the token and verify its validity
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": token.Valid},
		})
	}
}
