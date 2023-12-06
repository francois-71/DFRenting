package controllers

import (
    "context"
    "DFHousing/DFHousingAPI/configs"
    "DFHousing/DFHousingAPI/models"
    "DFHousing/DFHousingAPI/responses"
	"DFHousing/DFHousingAPI/utils/token"
	validator "DFHousing/DFHousingAPI/controllers/validators"
	"DFHousing/DFHousingAPI/controllers/validators"
    "net/http"
    "time"
	"encoding/base64"
	"io/ioutil"
	"fmt"
	"io"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// controller for property here.
// we can create, get, delete, update and get property here.
var propertyCollection *mongo.Collection = configs.GetCollection(configs.DB, "properties")
var fs *gridfs.Bucket = configs.GridFS

func CreateProperty() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.TODO()

		var property models.Property

		// Parse form data to get property details
		if err := c.ShouldBind(&property); err != nil {
			c.JSON(http.StatusBadRequest, responses.PropertyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		requiredFields := []string{
			"propertyname", "type", "description", "price_per_night",
			"number_of_bedrooms", "number_of_bathrooms", "house_rules",
			"cancellation_policy", "location", "city", "state", "zip", "country",
		}

		var bad_values []string
		// Validate required fields
		for _, field := range requiredFields {
			if len(c.PostForm(field)) == 0 {
				bad_values = append(bad_values, field)
			}
		}
		if len(bad_values) > 0 {
			c.JSON(http.StatusBadRequest, responses.PropertyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": fmt.Sprintf("Missing required fields: %v", bad_values)}})
			return
		}

		// Extract user ID from the token
		userIDFromToken, tokenErr := token.ExtractTokenID(c)
		if tokenErr != nil {
			c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": tokenErr.Error()}})
			return
		}

		// Fetch user details from the database
		realUserIDFromToken, _ := primitive.ObjectIDFromHex(userIDFromToken)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"id": realUserIDFromToken}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		approval := false
		// set approval to false if the user is not an admin and true if the user is an admin
		if user.Role == "admin" {
			approval = true
		}


		currentDateTime := time.Now()
		formattedDate := currentDateTime.Format("2006-01-02 15:04:05")

		// Set property fields from form data and other sources
		property = models.Property{
			Id:                primitive.NewObjectID(),
			HostID:            user.Id,
			HostName: 		   user.First_Name,
			PropertyName:      c.PostForm("propertyname"),
			Type:              c.PostForm("type"),
			Description:       c.PostForm("description"),
			PricePerNight:     c.PostForm("price_per_night"),
			NumberOfBedrooms:  c.PostForm("number_of_bedrooms"),
			NumberOfBathrooms: c.PostForm("number_of_bathrooms"),
			HouseRules:        c.PostForm("house_rules"),
			CancellationPolicy: c.PostForm("cancellation_policy"),
			Location:          c.PostForm("location"),
			City:              c.PostForm("city"),
			State:             c.PostForm("state"),
			Zip:               c.PostForm("zip"),
			Country:           c.PostForm("country"),
			Date:              formattedDate,
			Approval:          approval,
			IsActive:          true,
		}

		fmt.Printf("property: %v\n", property)
		if err := validator.Validate.Struct(property); err != nil {
			c.JSON(http.StatusBadRequest, responses.PropertyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.PropertyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer file.Close()
		contentType := header.Header.Get("Content-Type")

		// Create an http.Header and set the content type
		httpHeader := make(http.Header)
		httpHeader.Set("Content-Type", contentType)

		if err := validators.IsValidImageFormat(httpHeader); err != nil {
			c.JSON(http.StatusBadRequest, responses.PropertyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Use GridFS to store the image in MongoDB
		uploadStream, err := fs.OpenUploadStream(header.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer uploadStream.Close()

		_, err = io.Copy(uploadStream, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Update the property model with image information
		property.Image = models.Image{
			ImageURL:    header.Filename, // Store the filename as the image URL for now
			Filename:    header.Filename, // Save the filename for reference
			ContentType: header.Header.Get("Content-Type"),
			// Add other image-related fields if needed
		}

		// Save the property (including image info) in MongoDB
		result, err := propertyCollection.InsertOne(ctx, property)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.PropertyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}


func GetAProperty() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		propertyID := c.Param("propertyId")
		var property models.Property

		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(propertyID)

		err := propertyCollection.FindOne(ctx, bson.M{"id": objID}).Decode(&property)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		// Fetch image info if needed
		imageData, err := getImageData(property.Image.Filename)
		if err != nil {
			// Handle error fetching image
			property.Image.ImageURL = "" // Set empty URL or default image URL
		} else {
			property.Image.ImageURL = fmt.Sprintf("data:%s;base64,%s", property.Image.ContentType, base64.StdEncoding.EncodeToString(imageData))
		}

		// check approval status
		

		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": property},
		})
	}
}

func GetAllProperties() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cursor, err := propertyCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		defer cursor.Close(ctx)

		var properties []models.Property
		for cursor.Next(ctx) {
			var property models.Property
			if err := cursor.Decode(&property); err != nil {
				c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}

			// Check approval status, this way we return only approved properties
			if !property.Approval {
				continue
			}

			// Fetch image info if needed
			imageData, err := getImageData(property.Image.Filename)
			if err != nil {
				// Handle error fetching image
				property.Image.ImageURL = "" // Set empty URL or default image URL
			} else {
				property.Image.ImageURL = fmt.Sprintf("data:%s;base64,%s", property.Image.ContentType, base64.StdEncoding.EncodeToString(imageData))
			}

			properties = append(properties, property)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		if len(properties) == 0 {
			c.JSON(http.StatusNotFound, responses.PropertyResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    map[string]interface{}{"data": "No properties found"},
			})
			return
		}


		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": properties},
		})
	}
}

func GetPropertiesByFalseApproval() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := propertyCollection.Find(ctx, bson.M{"approval": false})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		defer cursor.Close(ctx)

		var properties []models.Property
		for cursor.Next(ctx) {
			var property models.Property
			if err := cursor.Decode(&property); err != nil {
				c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}

			// Fetch image info if needed
			imageData, err := getImageData(property.Image.Filename)
			if err != nil {
				// Handle error fetching image
				property.Image.ImageURL = "" // Set empty URL or default image URL
			} else {
				property.Image.ImageURL = fmt.Sprintf("data:%s;base64,%s", property.Image.ContentType, base64.StdEncoding.EncodeToString(imageData))
			}

			properties = append(properties, property)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		if len(properties) == 0 {
			c.JSON(http.StatusNotFound, responses.PropertyResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    map[string]interface{}{"data": "No properties found"},
			})
			return
		}

		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": properties},
		})
	
	}
}

func SetPropertyTrueApproval() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		propertyID := c.Param("propertyId")
		var property models.Property
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(propertyID)

		err := propertyCollection.FindOneAndUpdate(ctx, bson.M{"id": objID}, bson.M{"$set": bson.M{"approval": true}}).Decode(&property)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": property},
		})
	}
}

func SetPropertyFalseApproval() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		propertyID := c.Param("propertyId")
		var property models.Property
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(propertyID)

		//delete the property from the database
		err := propertyCollection.FindOneAndDelete(ctx, bson.M{"id": objID}).Decode(&property)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": property},
		})
	}
}


func getImageData(filename string) ([]byte, error) {
	image, err := fs.OpenDownloadStreamByName(filename)
	if err != nil {
		return nil, err
	}
	defer image.Close()

	imageData, err := ioutil.ReadAll(image)
	if err != nil {
		return nil, err
	}
	return imageData, nil
}

/*

func EditAProperty() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		propertyID := c.Param("propertyId")
		var property models.Property
		defer cancel()

		if err := c.Bind(&property); err != nil {
			c.JSON(http.StatusBadRequest, responses.PropertyResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Validate required fields if necessary
		// ...

		objID, _ := primitive.ObjectIDFromHex(propertyID)



		err := propertyCollection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update).Decode(&property)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": property},
		})
	}
}

func DeleteAProperty() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		propertyID := c.Param("propertyId")
		var property models.Property
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(propertyID)

		err := propertyCollection.FindOneAndDelete(ctx, bson.M{"_id": objID}).Decode(&property)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Delete associated image if necessary
		// ...

		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": property},
		})
	}
}



func GetPropertiesByHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hostID := c.Param("hostId")
		var properties []models.Property
		defer cancel()

		cursor, err := propertyCollection.Find(ctx, bson.M{"hostid": hostID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PropertyResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var property models.Property
			cursor.Decode(&property)

			//fetch and append image
			properties = append(properties, property)
		}

		c.JSON(http.StatusOK, responses.PropertyResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": properties},
		})
	}
}
*/