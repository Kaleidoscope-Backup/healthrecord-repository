package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// CodeSystemType ...
type CodeSystemType string

const (
	//INTERNAL ..
	INTERNAL CodeSystemType = "INTERNAL"

	//FHIR ..
	FHIR CodeSystemType = "FHIR"

	//ICD11 ..
	ICD11 CodeSystemType = "ICD11"

	//ICD10 ..
	ICD10 CodeSystemType = "ICD10"

	//ICD9 ..
	ICD9 CodeSystemType = "ICD9"

	//SNOMEDCT ..
	SNOMEDCT CodeSystemType = "SNOMEDCT"

	//LOINC ..
	LOINC CodeSystemType = "LOINC"

	//UMLS ..
	UMLS CodeSystemType = "UMLS"

	//CPT ..
	CPT CodeSystemType = "CPT"

	//RXNORM ..
	RXNORM CodeSystemType = "RXNORM"
)

// ConceptClassesInput ...
type ConceptClassesInput struct {
	ConceptClasses *[]ConceptClassInput `json:"conceptClasses"`
}

// ConceptClassQueryParam ...
type ConceptClassQueryParam struct {
	ExternalID *string `json:"externalID"`
	Name       *string `json:"name"`
	All        *bool   `json:"all"`
}

// ConceptClassInput ...
type ConceptClassInput struct {
	ExternalID  string    `json:"externalID"`
	Name        string    `json:"name"`
	Description TextInput `json:"description"`
}

// ConceptClass ...
type ConceptClass struct {
	Id          string       `json:"id" bson:"_id"`
	ExternalID  string       `json:"externalID" bson:"externalID"`
	Name        string       `json:"name" bson:"name"`
	Description Text         `json:"description" bson:"description"`
	Meta        *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// CodableConceptInput ...
type CodableConceptInput struct {
	Text         *string      `json:"text"`
	ConceptClass string       `json:"conceptClass"`
	Coding       *[]CodeInput `json:"code"`
}

// CodableConcept ...
type CodableConcept struct {
	Id           string       `json:"id" bson:"_id"`
	Text         *string      `json:"text" bson:"text"`
	ConceptClass string       `json:"conceptClass" bson:"conceptClass"`
	Coding       *[]Code      `json:"code" bson:"code"`
	Meta         *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// CodeInput ...
type CodeInput struct {
	System       CodeSystemType `json:"system"`
	Version      *string        `json:"version"`
	Code         string         `json:"code"`
	Display      string         `json:"display"`
	Definition   *string        `json:"definition"`
	Comment      *string        `json:"comment"`
	Language     *Language      `json:"language"`
	UserSelected *bool          `json:"userSelected"`
}

// Code ...
type Code struct {
	Id           string         `json:"id" bson:"_id"`
	System       CodeSystemType `json:"system" bson:"system"`
	Version      *string        `json:"version" bson:"version"`
	Code         string         `json:"code" bson:"code"`
	Display      string         `json:"display" bson:"display"`
	Definition   *string        `json:"definition" bson:"definition"`
	Comment      *string        `json:"comment" bson:"comment"`
	Language     *Language      `json:"language" bson:"language"`
	UserSelected *bool          `json:"userSelected" bson:"userSelected"`
	Meta         *models.Meta   //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// ClinicalCodeInput ...
type ClinicalCodeInput struct {
	Code       string         `json:"code"`
	Display    string         `json:"display"`
	Definition string         `json:"definition"`
	SystemType CodeSystemType `json:"systemType"`
	Language   *string        `json:"language"`
}

// ClinicalCode ...
type ClinicalCode struct {
	Id         string         `json:"id" bson:"_id"`
	SystemType CodeSystemType `json:"systemType" bson:"systemType"`
	Version    *string        `json:"version" bson:"version"`
	Code       string         `json:"code" bson:"code"`
	Display    string         `json:"display" bson:"display"`
	Definition string         `json:"definition" bson:"definition"`
	Language   *string        `json:"language" bson:"language"`
	Meta       *models.Meta   //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
