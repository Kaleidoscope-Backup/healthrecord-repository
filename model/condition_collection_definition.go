package model

import "github.com/karte/mongo-lib/models"

// ConditionTypeInput ...
type ConditionTypeInput struct {
	Name string               `json:"name"`
	Code *CodableConceptInput `json:"code"`
}

// ConditionType ...
type ConditionType struct {
	Id   string          `json:"id" bson:"_id"`
	Name string          `json:"name" bson:"name"`
	Code *CodableConcept `json:"code" bson:"code"`
	Meta *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// ConditionDefinitionCollectionInput ...
type ConditionDefinitionCollectionInput struct {
	Name       string                `json:"name"`
	Code       *CodableConceptInput  `json:"code"`
	Source     *string               `json:"source"`
	Language   Language              `json:"language"`
	Conditions *[]ConditionTypeInput `json:"conditions"`
}

// ConditionDefinitionCollection ...
type ConditionDefinitionCollection struct {
	Id         string           `json:"id" bson:"_id"`
	Name       string           `json:"name" bson:"name"`
	Code       *CodableConcept  `json:"code" bson:"code"`
	Source     *string          `json:"source" bson:"source"`
	Language   Language         `json:"language" bson:"language"`
	Conditions *[]ConditionType `json:"conditions" bson:"conditions"`
	Meta       *models.Meta     //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
