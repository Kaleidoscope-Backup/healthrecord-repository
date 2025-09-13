package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// ContactPointInput ...
type ContactPointInput struct {
	System string     `json:"system"`
	Value  string     `json:"value"`
	Rank   *int32     `json:"rank"`
	Start  *util.Time `json:"start"`
	End    *util.Time `json:"end"`
}

// ContactPoint ...
type ContactPoint struct {
	Id     string       `json:"id" bson:"_id"`
	System string       `json:"system" bson:"system"`
	Value  string       `json:"value" bson:"value"`
	Rank   *int32       `json:"rank" bson:"rank"`
	Start  *util.Time   `json:"start" bson:"start"`
	End    *util.Time   `json:"end" bson:"end"`
	Meta   *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
