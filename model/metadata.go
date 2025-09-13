package model

import "gitlab.com/karte/mongo-lib/models"

//AttributeInput ...
type AttributeInput struct {
	Name  string     `json:"name"`
	Value ValueInput `json:"value"`
}

//Attribute ...
type Attribute struct {
	Id    string       `json:"id" bson:"_id"`
	Name  string       `json:"name" bson:"name"`
	Value Value        `json:"value" bson:"value"`
	Meta  *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//MetaDataInput ...
type MetaDataInput struct {
	Name       string           `json:"name"`
	Value      string           `json:"value"`
	Attributes *[]MetaDataInput `json:"attributes"`
}

//MetaData ...
type MetaData struct {
	Id         string       `json:"id" bson:"_id"`
	Name       string       `json:"name" bson:"name"`
	Value      string       `json:"value" bson:"value"`
	Attributes *[]MetaData  `json:"attributes" bson:"attributes"`
	Meta       *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
