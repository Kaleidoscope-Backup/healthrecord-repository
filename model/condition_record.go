package model

import "github.com/karte/mongo-lib/models"

// ConditionStatus ...
type ConditionStatus string

const (
	//CONDITION_ACTIVE ..
	CONDITION_ACTIVE ConditionStatus = "CONDITION_ACTIVE"

	//CONDITION_RECURRENCE ..
	CONDITION_RECURRENCE ConditionStatus = "CONDITION_RECURRENCE"

	//CONDITION_INACTIVE ..
	CONDITION_INACTIVE ConditionStatus = "CONDITION_INACTIVE"

	//CONDITION_REMISSION ..
	CONDITION_REMISSION ConditionStatus = "CONDITION_REMISSION"

	//CONDITION_RESOLVED ..
	CONDITION_RESOLVED ConditionStatus = "CONDITION_RESOLVED"
)

// ConditionRecordCreate ...
type ConditionRecordCreate struct {
	HealthRecordCreate
	Code               *CodableConceptInput `json:"code"`
	Status             ConditionStatus      `json:"conditionStatus"`
	Severity           *Severity            `json:"severity"`
	BodySite           *string              `json:"bodySite"`
	BodySiteCode       *CodableConceptInput `json:"bodySiteCode"`
	StageAssesment     *string              `json:"stageAssesment"`
	StageAssesmentCode *CodableConceptInput `json:"stageAssesmentCode"`
	Onset              *OnsetInput          `json:"onset"`
	Abatement          *AbatementInput      `json:"abatement"`
	Evidence           *[]SymptomInput      `json:"evidence" bson:"evidene"`
}

// ConditionRecord ...
type ConditionRecord struct {
	HealthRecord
	Id                 string          `json:"id" bson:"_id"`
	Code               *CodableConcept `json:"code" bson:"code"`
	Status             ConditionStatus `json:"status" bson:"status"`
	Severity           *Severity       `json:"severity" bson:"severity"`
	BodySite           *string         `json:"bodySite" bson:"bodySite"`
	BodySiteCode       *CodableConcept `json:"bodySiteCode" bson:"bodySiteCode"`
	StageAssesment     *string         `json:"stageAssesment" bson:"stageAssesment"`
	StageAssesmentCode *CodableConcept `json:"stageAssesmentCode" bson:"stageAssesmentCode"`
	Onset              *Onset          `json:"onset" bson:"onset"`
	Abatement          *Abatement      `json:"abatement" bson:"abatement"`
	Evidence           *[]Symptom      `json:"evidence" bson:"evidene"`
	Meta               *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
