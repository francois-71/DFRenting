package controllers

import (
    "context"
    "DFHousing/DFHousingAPI/configs"
    "DFHousing/DFHousingAPI/models"
    "DFHousing/DFHousingAPI/responses"
	validator "DFHousing/DFHousingAPI/controllers/validators"
	
    "net/http"
    "time"


    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var reviewCollection *mongo.Collection = configs.GetCollection(configs.DB, "reviews")

func CreateReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var review models.Review
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validator.Validate.Struct(&review); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newReview := models.Review{
			Id:       primitive.NewObjectID(),
			PropertyID: review.PropertyID,
			UserID:      review.UserID,
			Rating:  review.Rating,
			Review:   review.Review,
		}
	  
		result, err := userCollection.InsertOne(ctx, newReview)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ReviewResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.ReviewResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var review models.Review

		//validate the request body
		if err := c.BindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validator.Validate.Struct(&review); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		
		filter := bson.M{"_id": review.Id}
		err := userCollection.FindOne(ctx, filter).Decode(&review)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ReviewResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ReviewResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": review}})
	}
}

func EditAReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var review models.Review
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validator.Validate.Struct(&review); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		
		filter := bson.M{"_id": review.Id}
		update := bson.M{
			"$set": bson.M{
				"rating": review.Rating,
				"review": review.Review,
			},
		}
		result, err := userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ReviewResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ReviewResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DeleteAReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var review models.Review
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validator.Validate.Struct(&review); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		
		filter := bson.M{"_id": review.Id}
		result, err := userCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ReviewResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ReviewResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetReviewsByPropertyId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var review models.Review
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validator.Validate.Struct(&review); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ReviewResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		
		filter := bson.M{"propertyid": review.PropertyID}
		cursor, err := userCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ReviewResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var reviews []bson.M
		if err = cursor.All(ctx, &reviews); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ReviewResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ReviewResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": reviews}})
	}
}




