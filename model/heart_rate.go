package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

type HeartRate struct {
	Id    string       `json:"id" bson:"_id"`
	Value int32        `json:"value" bson:"value"`
	Unit  string       `json:"unit" bson:"unit"`
	Meta  *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
