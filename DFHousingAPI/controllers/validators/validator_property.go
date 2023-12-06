package validators

// validator for the validator controller
// checks rules for HOSTID, HOSTNAME, NAME, TYPE, DESCRIPTION, PRICE_PER_NIGHT, NUMBER_OF_BEDROOMS, NUMBER_OF_BATHROOMS, HOUSE_RULES, CANCELLATION_POLICY, LOCATION, CITY, STATE, ZIP, COUNTRY, DATE, REVIEWS, IMAGE, IS_ACTIVE

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"fmt"
	"strconv"
)
/*
package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
    Id                 primitive.ObjectID `json:"id,omitempty"`
    HostID             primitive.ObjectID `json:"hostid,omitempty" bson:"hostid,omitempty"`
	HostName		   string             `json:"hostname,omitempty" bson:"hostname,omitempty"`
    Type               string             `json:"type,omitempty" validate:"required"`
    Description        string             `json:"description,omitempty" validate:"required"`
    PricePerNight      string             `json:"price_per_night,omitempty" validate:"required"`
    NumberOfBedrooms   string             `json:"number_of_bedrooms,omitempty" validate:"required"`
    NumberOfBathrooms  string             `json:"number_of_bathrooms,omitempty" validate:"required"`
    HouseRules         string             `json:"house_rules,omitempty" validate:"required"`
    CancellationPolicy string             `json:"cancellation_policy,omitempty" validate:"required"`
    Location           string             `json:"location,omitempty" validate:"required"`
    City               string             `json:"city,omitempty" validate:"required"`
    State              string             `json:"state,omitempty" validate:"required"`
    Zip                string             `json:"zip,omitempty" validate:"required"`
    Country            string             `json:"country,omitempty" validate:"required"`
	Date 			   string             `json:"date,omitempty"`
	Reviews			   []Review           `json:"reviews,omitempty"`
    Image              Image              `json:"image,omitempty" bson:"image_id,omitempty"`
    IsActive           bool               `json:"is_active,omitempty"`
}
*/

func CustomPropertyNameValidation(fl validator.FieldLevel) bool {
	propertyName := fl.Field().String()
	fmt.Printf("checking name: %s\n", propertyName)
	// Regular expression for validating hostname format
	propertyNameRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the hostname length is not more than 100 characters and not less than 6 characters
	if len(propertyName) < 6 {
		return false
	}

	if len(propertyName) > 100 {
		return false
	}

	// Check if the name matches the regex pattern
	if !propertyNameRegex.MatchString(propertyName) {
		return false
	}

	return true
}

func CustomTypeValidation(fl validator.FieldLevel) bool {
	propertyType := fl.Field().String()
	// set propertyType to lowercase
	propertyType = strings.ToLower(propertyType)
	fmt.Printf("checking type: %s\n", propertyType)
	
	// check if the property type is not empty
	if propertyType == "" {
		return false
	}

	// Property type can only be flat or house
	if propertyType != "flat" && propertyType != "house" {
		return false
	}

	return true
}

func CustomDescriptionValidation(fl validator.FieldLevel) bool {
	description := fl.Field().String()
	fmt.Printf("checking description: %s\n", description)
	// Regular expression for validating description format
	descriptionRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the description length is not more than 250 characters
	if len(description) < 6 {
		return false
	}

	if len(description) > 250 {
		return false
	}

	// Check if the description matches the regex pattern
	if !descriptionRegex.MatchString(description) {
		return false
	}

	return true
}

func CustomPricePerNightValidation(fl validator.FieldLevel) bool {
	pricePerNight := fl.Field().String()

	// Convert the price to integer
	pricePerNightInt, err := strconv.Atoi(pricePerNight)
	if err != nil {
		return false
	}

	// Check if the price_per_night is not less than 0
	if pricePerNightInt < 0 {
		return false
	}

	// Check if the price_per_night is not more than 100000
	if pricePerNightInt > 100000 {
		return false
	}

	// Check if the price_per_night is not empty
	if pricePerNight == "" {
		return false
	}

	return true
}

func CustomNumberOfBedroomsValidation(fl validator.FieldLevel) bool {
	numberOfBedrooms := fl.Field().String()
	fmt.Printf("checking number_of_bedrooms: %s\n", numberOfBedrooms)
	numberOfBedroomsInt, err := strconv.Atoi(numberOfBedrooms)
	if err != nil {
		return false
	}

	// Check if the number_of_bedrooms is not less than 0
	if numberOfBedroomsInt < 0 {
		return false
	}

	// Check if the number_of_bedrooms is not more than 100
	if numberOfBedroomsInt > 100 {
		return false
	}

	// Check if the number_of_bedrooms is not empty
	if numberOfBedrooms == "" {
		return false
	}

	return true
}

func CustomNumberOfBathroomsValidation(fl validator.FieldLevel) bool {
	numberOfBathrooms := fl.Field().String()
	
	numberOfBathroomsInt, err := strconv.Atoi(numberOfBathrooms)
	if err != nil {
		return false
	}
	
	// Check if the number_of_bathrooms is not less than 0
	if numberOfBathroomsInt < 0 {
		return false
	}

	// Check if the number_of_bathrooms is not more than 100
	if numberOfBathroomsInt > 100 {
		return false
	}

	// Check if the number_of_bathrooms is not empty
	if numberOfBathrooms == "" {
		return false
	}

	return true
}

func CustomHouseRulesValidation(fl validator.FieldLevel) bool {
	houseRules := fl.Field().String()
	fmt.Printf("checking house_rules: %s\n", houseRules)
	// Regular expression for validating house_rules format
	houseRulesRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the house_rules length is not more than 250 characters

	if len(houseRules) > 250 {
		return false
	}

	// Check if the house_rules matches the regex pattern
	if !houseRulesRegex.MatchString(houseRules) {
		return false
	}

	return true
}

func CustomCancellationPolicyValidation(fl validator.FieldLevel) bool {
	cancellationPolicy := fl.Field().String()
	fmt.Printf("checking cancellation_policy: %s\n", cancellationPolicy)
	// Regular expression for validating cancellation_policy format
	cancellationPolicyRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the cancellation_policy length is not more than 250 characters

	if len(cancellationPolicy) > 250 {
		return false
	}

	// Check if the cancellation_policy matches the regex pattern
	if !cancellationPolicyRegex.MatchString(cancellationPolicy) {
		return false
	}

	return true
}

func CustomLocationValidation(fl validator.FieldLevel) bool {
	location := fl.Field().String()
	fmt.Printf("checking location: %s\n", location)
	// Regular expression for validating location format
	locationRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the location length is not more than 250 characters
	if len(location) < 1 {
		return false
	}

	if len(location) > 250 {
		return false
	}

	// Check if the location matches the regex pattern
	if !locationRegex.MatchString(location) {
		return false
	}

	return true
}

func CustomCityValidation(fl validator.FieldLevel) bool {
	city := fl.Field().String()
	fmt.Printf("checking city: %s\n", city)
	// Regular expression for validating city format
	cityRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)


	// Check if the city length is not more than 250 characters 
	if len(city) < 1 {
		return false
	}

	if len(city) > 250 {
		return false
	}

	// Check if the city matches the regex pattern
	if !cityRegex.MatchString(city) {
		return false
	}

	return true
}

func CustomStateValidation(fl validator.FieldLevel) bool {
	state := fl.Field().String()
	fmt.Printf("checking state: %s\n", state)
	// Regular expression for validating state format
	stateRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the state length is not more than 250 characters
	if len(state) < 1 {
		return false
	}

	if len(state) > 250 {
		return false
	}

	// Check if the state matches the regex pattern
	if !stateRegex.MatchString(state) {
		return false
	}

	return true
}

func CustomZipValidation(fl validator.FieldLevel) bool {
	zip := fl.Field().String()
	fmt.Printf("checking zip: %s\n", zip)
	// Regular expression for validating zip format
	zipRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the zip length is not more than 250 characters
	if len(zip) < 1 {
		return false
	}

	if len(zip) > 250 {
		return false
	}

	// Check if the zip matches the regex pattern
	if !zipRegex.MatchString(zip) {
		return false
	}

	return true
}

func CustomCountryValidation(fl validator.FieldLevel) bool {
	country := fl.Field().String()
	fmt.Printf("checking country: %s\n", country)
	// Regular expression for validating country format
	countryRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-\s]+$`)

	// Check if the country length is not more than 250 characters
	if len(country) < 1 {
		return false
	}

	if len(country) > 250 {
		return false
	}

	// Check if the country matches the regex pattern
	if !countryRegex.MatchString(country) {
		return false
	}

	return true
}

/*
func CustomReviewsValidation(fl validator.FieldLevel) bool {
	reviews := fl.Field().String()
	fmt.Printf("checking reviews: %s\n", reviews)
	// Regular expression for validating reviews format
	reviewsRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+$`)

	// Check if the reviews length is not more than 250 characters
	if len(reviews) < 6 {
		return false
	}

	if len(reviews) > 250 {
		return false
	}

	// Check if the reviews matches the regex pattern
	if !reviewsRegex.MatchString(reviews) {
		return false
	}

	return true
}
*/



