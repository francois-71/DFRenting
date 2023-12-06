package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
    Id                 primitive.ObjectID `json:"id,omitempty"`
    HostID             primitive.ObjectID `json:"hostid,omitempty" bson:"hostid,omitempty"`
    HostName           string             `json:"hostname,omitempty" bson:"hostname,omitempty"`
    PropertyName       string             `json:"propertyname,omitempty" validate:"required,propertyname" bson:"propertyname,omitempty"`
    Type               string             `json:"type,omitempty" validate:"required,type"`
    Description        string             `json:"description,omitempty" validate:"required,description"`
    PricePerNight      string             `json:"price_per_night,omitempty" validate:"required,price_per_night"`
    NumberOfBedrooms   string             `json:"number_of_bedrooms,omitempty" validate:"required,number_of_bedrooms"`
    NumberOfBathrooms  string             `json:"number_of_bathrooms,omitempty" validate:"required,number_of_bathrooms"`
    HouseRules         string             `json:"house_rules,omitempty" validate:"required,house_rules"`
    CancellationPolicy string             `json:"cancellation_policy,omitempty" validate:"required,cancellation_policy"`
    Location           string             `json:"location,omitempty" validate:"required,location"`
    City               string             `json:"city,omitempty" validate:"required,city"`
    State              string             `json:"state,omitempty" validate:"required,state"`
    Zip                string             `json:"zip,omitempty" validate:"required,zip"`
    Country            string             `json:"country,omitempty" validate:"required,country"`
	Date 			   string             `json:"date,omitempty"`
	Reviews			   []Review           `json:"reviews,omitempty"`
    Approval           bool               `json:"approval,omitempty"`
    Image              Image              `json:"image,omitempty" bson:"image_id,omitempty"`
    IsActive           bool               `json:"is_active,omitempty"`
}