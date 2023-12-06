package controllers 

import (
	"context"
	"DFHousing/DFHousingAPI/models"
	"DFHousing/DFHousingAPI/responses"
	validator "DFHousing/DFHousingAPI/controllers/validators"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		
		//use the validator library to validate required fields
		if validationErr := validator.Validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		
		// check if user already exists
		filter := bson.M{"email": user.Email}
		err := userCollection.FindOne(ctx, filter).Decode(&user)
		if err == nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "An account with this email already exists"}})
			return
		}

		// hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			First_Name: user.First_Name,
			Email:      user.Email,
			Last_Name:  user.Last_Name,
			Password:   string(hashedPassword),
			Phone:      user.Phone,
			Age:        user.Age,
			Address:    user.Address,
			City:       user.City,
			State:      user.State,
			Zip:        user.Zip,
			Country:    user.Country,
			Role:       "user",
			IsActive:   true,
		}
	  
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}