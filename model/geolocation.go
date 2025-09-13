package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// GeoLocation represents a specific coordinate
type GeoLocation struct {
	Id        string       `json:"id" bson:"_id"`
	Name      *string      `json:"name" bson:"name"`
	Latitude  float64      `json:"latitude" bson:"latitude"`
	Longitude float64      `json:"longitude" bson:"longitude"`
	Elevation *float64     `json:"elevation" bson:"elevation"`
	Meta      *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// GeoLocationInput represents a specific coordinate
type GeoLocationInput struct {
	Name      *string  `json:"name"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Elevation *float64 `json:"elevation" bson:"elevation"`
}
