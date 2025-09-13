package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// SourceRecordIDInput ...
type SourceRecordIDInput struct {
	System string `json:"system"`
	Value  string `json:"qualification"`
}

// SourceRecordID ....
type SourceRecordID struct {
	Id     string       `json:"id" bson:"_id"`
	System string       `json:"system" bson:"system"`
	Value  string       `json:"qualification" bson:"qualification"`
	Meta   *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
