package model

import "github.com/karte/mongo-lib/models"

// AddressInput ...
type AddressInput struct {
	Name         *string `json:"name" bson:"name"`
	StreetNumber string  `json:"streetNumber"`
	StreetName   string  `json:"streetName"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	ZipCode      string  `json:"zipCode,omitempty"`
}

// Address represents any location
type Address struct {
	Id           string       `json:"id" bson:"_id"`
	Name         *string      `json:"name" bson:"name"`
	StreetNumber string       `json:"streetNumber" bson:"streetNumber"`
	StreetName   string       `json:"streetName" bson:"streetName"`
	City         string       `json:"city" bson:"city"`
	State        string       `json:"state" bson:"state"`
	Country      string       `json:"country" bson:"country"`
	ZipCode      string       `json:"zipCode,omitempty" bson:"zipCode,omitempty"`
	Location     *GeoLocation `json:"location,omitempty" bson:"location,omitempty"`
	Meta         *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
