package model

import "github.com/karte/mongo-lib/models"

// ObservationDefinitionCollectionQueryParam ...
type ObservationDefinitionCollectionQueryParam struct {
	Name      *string   `json:"name"`
	Publisher *string   `json:"publisher"`
	Language  *Language `json:"language"`
}

// MeasurementDefinitionInput ...
type MeasurementDefinitionInput struct {
	Name            string                 `json:"name"`
	Unit            string                 `json:"unit"`
	LowerLimit      int32                  `json:"lowerLimit"`
	UpperLimit      int32                  `json:"upperLimit"`
	Code            *CodableConceptInput   `json:"code"`
	ObservationType *HealthRecordType      `json:"observationType"`
	Attributes      *[]AttributeInput      `json:"attributes"`
	ReferenceRanges *[]ReferenceRangeInput `json:"referenceRanges"`
}

// MeasurementDefinition ...
type MeasurementDefinition struct {
	Id              string            `json:"id" bson:"_id"`
	Name            string            `json:"name" bson:"name"`
	Unit            string            `json:"unit" bson:"unit"`
	LowerLimit      int32             `json:"lowerLimit" bson:"lowerLimit"`
	UpperLimit      int32             `json:"upperLimit" bson:"upperLimit"`
	Code            *CodableConcept   `json:"code" bson:"code"`
	ObservationType *HealthRecordType `json:"observationType" bson:"observationType"`
	Attributes      *[]Attribute      `json:"attributes" bson:"attributes"`
	ReferenceRanges *[]ReferenceRange `json:"referenceRanges" bson:"referenceRanges"`
}

// ObservationDefinitionCollectionInput ...
type ObservationDefinitionCollectionInput struct {
	Name         string                        `json:"name"`
	Purpose      *string                       `json:"purpose"`
	Description  *string                       `json:"description"`
	Publisher    *string                       `json:"publisher"`
	Source       *SourceInput                  `json:"source"`
	Language     Language                      `json:"language"`
	Code         *CodableConceptInput          `json:"code"`
	Attributes   *[]AttributeInput             `json:"attributes"`
	Measurements *[]MeasurementDefinitionInput `json:"measurements"`
}

// ObservationDefinitionCollection ...
type ObservationDefinitionCollection struct {
	Id           string                   `json:"id" bson:"_id"`
	Name         string                   `json:"name" bson:"name"`
	Purpose      *string                  `json:"purpose" bson:"purpose"`
	Description  *string                  `json:"description" bson:"description"`
	Publisher    *string                  `json:"publisher" bson:"publisher"`
	Source       *Source                  `json:"source" bson:"source"`
	Language     Language                 `json:"language" bson:"language"`
	Code         *CodableConcept          `json:"code" bson:"code"`
	Attributes   *[]Attribute             `json:"attributes" bson:"attributes"`
	Measurements *[]MeasurementDefinition `json:"measurements" bson:"measurements"`
	Meta         *models.Meta             //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
