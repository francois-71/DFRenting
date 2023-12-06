package controllers

import (
    "context"
    "DFHousing/DFHousingAPI/configs"
    "DFHousing/DFHousingAPI/models"
    "DFHousing/DFHousingAPI/responses"
    "DFHousing/DFHousingAPI/utils/token"
    validator "DFHousing/DFHousingAPI/controllers/validators"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"

)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

//admin function later to create a user when connected as an admin
func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

        //Validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to Validate required fields
        if validationErr := validator.Validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        var existingUser models.User
        err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
        if err == nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Email already exists"}})
            return
        }

        newUser := models.User{
            Id:       primitive.NewObjectID(),
            First_Name: user.First_Name,
            Email:      user.Email,
            Last_Name:  user.Last_Name,
            Password:   user.Password,
            Phone:      user.Phone,
            Age:        user.Age,
            Address:    user.Address,
            City:       user.City,
            State:      user.State,
            Zip:        user.Zip,
            Country:    user.Country,
            Role:       user.Role,
            IsActive:   user.IsActive,
        }
      
        result, err := userCollection.InsertOne(ctx, newUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}

func CreateAdminUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

        //Validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to Validate required fields
        if validationErr := validator.Validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        //check if email already exists
        var existingUser models.User
        err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
        if err == nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Email already exists"}})
            return
        }

        newUser := models.User{
            Id:       primitive.NewObjectID(),
            First_Name: user.First_Name,
            Email:      user.Email,
            Last_Name:  user.Last_Name,
            Password:   user.Password,
            Phone:      user.Phone,
            Age:        user.Age,
            Address:    user.Address,
            City:       user.City,
            State:      user.State,
            Zip:        user.Zip,
            Country:    user.Country,
            Role:       "admin",
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

func IsUserAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        var user models.User

        //get the user id from the token
        userIDFromToken, tokenErr := token.ExtractTokenID(c)
        if tokenErr != nil {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": tokenErr.Error()}})
            return
        }

        //get the user role from the token
        userRoleFromToken, tokenErr := token.ExtractTokenRole(c)
        if tokenErr != nil {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": tokenErr.Error()}})
            return
        }

        //check the token role
        if string(userRoleFromToken) != "admin" {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "You are not authorized to perform this action"}})
            return
        }

        // we perform a second check to ensure that the user is an admin by looking in the database
        objId, _ := primitive.ObjectIDFromHex(userIDFromToken)

        err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //check the user role
        if user.Role != "admin" {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "You are not authorized to perform this action"}})
            return
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User is an admin"}})
    }   
}

// Function to get a user by id, admin can get any user, user can only get their own details
func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

        userId := c.Param("userId")

        //get the user id from the token
        userIDFromToken, tokenErr := token.ExtractTokenID(c)
        userRoleFromToken, tokenErr := token.ExtractTokenRole(c)
        if tokenErr != nil {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error4", Data: map[string]interface{}{"data": tokenErr.Error()}})
            return
        }

        //compare the user id from the token and the url to prevent a user to view another user's details

        if userId != string(userIDFromToken) && string(userRoleFromToken) != "admin" {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error5", Data: map[string]interface{}{"data": "You are not authorized to view this user"}})
            return
        }

        var user models.User
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

        err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
    }
}

func EditAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        var user models.User
        defer cancel()
        objId, _ := primitive.ObjectIDFromHex(userId)

        //Validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to Validate required fields
        if validationErr := validator.Validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        update := bson.M{
            "$set": bson.M{
                "first_name": user.First_Name,
                "email":      user.Email,
                "last_name":  user.Last_Name,
                "password":   user.Password,
                "phone":      user.Phone,
                "age":        user.Age,
                "address":    user.Address,
                "city":       user.City,
                "state":      user.State,
                "zip":        user.Zip,
                "country":    user.Country,
                "role":       user.Role,
                "isactive":   user.IsActive,
            },
        }

        result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //get updated user details
        var updatedUser models.User
        if result.MatchedCount == 1 {
            err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
    }
}

func DeleteAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

        result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound,
                responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
            )
            return
        }

        c.JSON(http.StatusOK,
            responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
        )
    }
}

func GetAllUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var users []models.User
        defer cancel()

        tokenRole, tokenErr := token.ExtractTokenRole(c)
        if tokenErr != nil {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": tokenErr.Error()}})
            return
        }
        if tokenRole != "admin" {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "You are not authorized to view all users"}})
            return
        }

        results, err := userCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleUser models.User
            if err = results.Decode(&singleUser); err != nil {
                c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
          
            users = append(users, singleUser)
        }
        for i := 0; i < len(users); i++ {
            users[i].Password = ""
        }

        c.JSON(http.StatusOK,
            responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
        )
    }
}

// get the user from the token
func GetCurrentUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        var user models.User
        
        userIDFromToken, tokenErr := token.ExtractTokenID(c)
        if tokenErr != nil {
            c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": tokenErr.Error()}})
            return
        }

        objId, _ := primitive.ObjectIDFromHex(userIDFromToken)

        err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
        //remove the password from the response
        user.Password = ""

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
    }
}
