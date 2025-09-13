package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// SourceInput ...
type SourceInput struct {
	Name        string  `json:"name"`
	URI         string  `json:"uri"`
	Description *string `json:"description"`
}

// Source ...
type Source struct {
	Id          string       `json:"id" bson:"_id"`
	Name        string       `json:"name" bson:"name"`
	URI         string       `json:"uri" bson:"uri"`
	Description *string      `json:"description" bson:"description"`
	Meta        *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
