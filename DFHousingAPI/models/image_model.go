package models 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	Id                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ImageURL		  string             `json:"imageurl,omitempty" bson:"imageurl,omitempty"`
	Filename 		string             `json:"filename,omitempty" bson:"filename,omitempty"`
	ContentType 	string             `json:"contentType,omitempty" bson:"contentType,omitempty"`
}