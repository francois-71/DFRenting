package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	PropertyID  primitive.ObjectID `json:"propertyid,omitempty"`
	UserID  primitive.ObjectID `json:"reviewuserid,omitempty"`
	Rating string `json:"rating,omitempty validate:"required"`
	Review string `json:"review,omitempty validate:"required"`
	Date string `json:"reviewdate,omitempty validate:"required"`
}