package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// ProcedureStatus ...
type ProcedureStatus string

const (
	//PROCEDURE_PREPARATION ..
	PROCEDURE_PREPARATION ProcedureStatus = "PROCEDURE_PREPARATION"
	//PROCEDURE_IN_PROGRESS ..
	PROCEDURE_IN_PROGRESS ProcedureStatus = "PROCEDURE_IN_PROGRESS"
	//PROCEDURE_SUSPENDED ..
	PROCEDURE_SUSPENDED ProcedureStatus = "PROCEDURE_SUSPENDED"
	//PROCEDURE_ABORTED ..
	PROCEDURE_ABORTED ProcedureStatus = "PROCEDURE_ABORTED"
	//PROCEDURE_COMPLETED ..
	PROCEDURE_COMPLETED ProcedureStatus = "PROCEDURE_COMPLETED"
	//PROCEDURE_COMPLETED ..
	PROCEDURE_ENTERED_IN_ERROR ProcedureStatus = "PROCEDURE_ENTERED_IN_ERROR"
	//PROCEDURE_UNKNOWN ..
	PROCEDURE_UNKNOWN ProcedureStatus = "PROCEDURE_UNKNOWN"
)

// ProcedureOutcome ...
type ProcedureOutcome string

const (
	//PROCEDURE_SUCCESSFUL ..
	PROCEDURE_SUCCESSFUL ProcedureOutcome = "PROCEDURE_SUCCESSFUL"
	//PROCEDURE_UN_SUCCESSFUL ..
	PROCEDURE_UN_SUCCESSFUL ProcedureOutcome = "PROCEDURE_UN_SUCCESSFUL"
	//PROCEDURE_PARTIALLY_SUCCESSFUL ..
	PROCEDURE_PARTIALLY_SUCCESSFUL ProcedureOutcome = "PROCEDURE_PARTIALLY_SUCCESSFUL"
)

// ProcedureCategory ...
type ProcedureCategory string

const (
	//PROCEDURE_PSYCHIATRY ..
	PROCEDURE_PSYCHIATRY ProcedureCategory = "PROCEDURE_PSYCHIATRY"
	//PROCEDURE_COUNSELLING ..
	PROCEDURE_COUNSELLING ProcedureCategory = "PROCEDURE_COUNSELLING"
	//PROCEDURE_SURGERY ..
	PROCEDURE_SURGERY ProcedureCategory = "PROCEDURE_SURGERY"
)

// ProcedureCode ...
type ProcedureCode string

const (
	CODE_24642003  ProcedureCode = "Psychiatry procedure or service"
	CODE_409063005 ProcedureCode = "Counselling"
	CODE_409073007 ProcedureCode = "Education"
	CODE_387713003 ProcedureCode = "Surgical procedure"
	CODE_103693007 ProcedureCode = "Diagnostic procedure"
	CODE_46947000  ProcedureCode = "Chiropractic manipulation"
)

// ProcedurRecordCreate ...
type ProcedurRecordCreate struct {
	HealthRecordCreate
	Status              ProcedureStatus      `json:"status"`
	Category            ProcedureCategory    `json:"category"`
	Code                *CodableConceptInput `json:"code"`
	Performer           *string              `json:"performer"`
	Reason              string               `json:"reason"`
	ReasonCode          *CodableConceptInput `json:"reasonCode"`
	BodySite            *string              `json:"bodySite"`
	BodySiteCode        *CodableConceptInput `json:"bodySiteCode"`
	Outcome             *ProcedureOutcome    `json:"outcome"`
	OutcomeCode         *CodableConceptInput `json:"outcomeCode"`
	FollowupInstruction *string              `json:"followupInstruction"`
	Report              *string              `json:"report"`
}

// ProcedureRecord ...
type ProcedureRecord struct {
	HealthRecord
	Id                  string            `json:"id" bson:"_id"`
	Status              ProcedureStatus   `json:"status" bson:"status"`
	Category            ProcedureCategory `json:"category" bson:"category"`
	Code                *CodableConcept   `json:"code" bson:"code"`
	Performer           *string           `json:"performer" bson:"performer"`
	Reason              string            `json:"reason" bson:"reason"`
	ReasonCode          *CodableConcept   `json:"reasonCode" bson:"reasonCode"`
	BodySite            *string           `json:"bodySite" bson:"bodySite"`
	BodySiteCode        *CodableConcept   `json:"bodySiteCode" bson:"bodySiteCode"`
	Outcome             *ProcedureOutcome `json:"outcome" bson:"outcome"`
	OutcomeCode         *CodableConcept   `json:"outcomeCode" bson:"outcomeCode"`
	FollowupInstruction *string           `json:"followupInstruction" bson:"followupInstruction"`
	Report              *string           `json:"report" bson:"report"`
	Meta                *models.Meta      //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection

}
