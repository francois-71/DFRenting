package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    Id         primitive.ObjectID `json:"id,omitempty"`
    First_Name string             `json:"first_name,omitempty" validate:"required"`
    Email      string             `json:"email,omitempty" validate:"required"`
    Last_Name  string             `json:"last_name,omitempty" validate:"required"`
    Password   string             `json:"password,omitempty" validate:"required"`
    Phone      string             `json:"phone,omitempty" validate:"required"`
    Age        string               `json:"age,omitempty" validate:"required"`
    Address    string             `json:"address,omitempty"`
    City       string             `json:"city,omitempty"`
    State      string             `json:"state,omitempty"`
    Zip        string             `json:"zip,omitempty"`
    Country    string             `json:"country,omitempty" validate:"required"`
    Role       string             `json:"role,omitempty"`
    IsActive   bool               `json:"isactive,omitempty"`
}
