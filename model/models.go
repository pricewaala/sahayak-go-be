package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive" 
)

type Workers struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` 
	Name   string             `json:"name,omitempty"`
	PhoneNumber string               `json:"phone,omitempty"`
	City string	`json:"city,omitempty"`
	State string `json:"state,omitempty"`
	PinCide string `json:"pincode,omitempty"`
}
