package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// ConsentAction ...
type ConsentAction string

const (
	//COLLECT ..
	COLLECT ConsentAction = "COLLECT"
	//ACCESS ..
	ACCESS ConsentAction = "ACCESS"
	//USE ..
	USE ConsentAction = "USE"
	//DISCLOSE ..
	DISCLOSE ConsentAction = "DISCLOSE"
	//CORRECT ..
	CORRECT ConsentAction = "CORRECT"
)

// PurposeOfUse ...
type PurposeOfUse string

const (
	//HEALTHCARE_MARKETING ..
	HEALTHCARE_MARKETING PurposeOfUse = "HEALTHCARE_MARKETING"

	//HEALTHCARE_OPERATIONS ..
	HEALTHCARE_OPERATIONS PurposeOfUse = "HEALTHCARE_OPERATIONS"

	//DONATION ..
	DONATION PurposeOfUse = "DONATION"

	//FRAUD ..
	FRAUD PurposeOfUse = "FRAUD"

	//GOVERNMENT_USE ..
	GOVERNMENT_USE PurposeOfUse = "GOVERNMENT_USE"

	//HEALTH_ACCREDITION ..
	HEALTH_ACCREDITION PurposeOfUse = "HEALTH_ACCREDITION"

	//DECEDENT ..
	DECEDENT PurposeOfUse = "DECEDENT"

	//LEGAL ..
	LEGAL PurposeOfUse = "LEGAL"

	//HEALTH_OUTCOME_MEASURE ..
	HEALTH_OUTCOME_MEASURE PurposeOfUse = "HEALTH_OUTCOME_MEASURE"

	//HEALTH_PROGRAM_REPORTING ..
	HEALTH_PROGRAM_REPORTING PurposeOfUse = "HEALTH_PROGRAM_REPORTING"

	//HEALTH_QUALITY_IMPROVEMENT ..
	HEALTH_QUALITY_IMPROVEMENT PurposeOfUse = "HEALTH_QUALITY_IMPROVEMENT"

	//HEALTHCARE_RESEARCH ..
	HEALTHCARE_RESEARCH PurposeOfUse = "HEALTHCARE_RESEARCH"

	//CLINICAL_TRIAL ..
	CLINICAL_TRIAL PurposeOfUse = "CLINICAL_TRIAL"

	//EMERGENCY_TREATMENT ..
	EMERGENCY_TREATMENT PurposeOfUse = "EMERGENCY_TREATMENT"
)

// ConsentCreate ...
type ConsentCreate struct {
	Name                  *string                `json:"name"`
	Content               *string                `json:"content"`
	Category              string                 `json:"category"`
	Context               *ReferenceEntity       `json:"context"`
	ConsumerID            string                 `json:"consumerID"`
	Period                *PeriodInput           `json:"period"`
	ConsentingParty       *[]ReferenceActorInput `json:"consentingParty"`
	Custodian             *string                `json:"custodian"`
	Action                *ConsentAction         `json:"action"`
	Purpose               PurposeOfUse           `json:"purpose"`
	QuestionnaireResponse *string                `json:"questionnaireResponse"`
	DateTime              util.Time              `json:"dateTime"`
}

// Consent ...
type Consent struct {
	Id                    string            `json:"id" bson:"_id"`
	Name                  *string           `json:"name" bson:"name"`
	Content               *string           `json:"content" bson:"content"`
	Category              string            `json:"category" bson:"category"`
	CategoryCode          *ClinicalCode     `json:"categoryCode" bson:"categoryCode"`
	Context               *ReferenceEntity  `json:"context" bson:"context"`
	ConsumerID            string            `json:"consumerID" bson:"consumerID"`
	Period                *Period           `json:"period" bson:"period"`
	ConsentingParty       *[]ReferenceActor `json:"consentingParty" bson:"consentingParty"`
	Custodian             *string           `json:"custodian" bson:"custodian"`
	Action                *ConsentAction    `json:"action" bson:"action"`
	Purpose               PurposeOfUse      `json:"purpose" bson:"purpose"`
	QuestionnaireResponse *string           `json:"questionnaireResponse" bson:"questionnaireResponse"`
	PurposeCode           *ClinicalCode     `json:"purposeCode" bson:"purposeCode"`
	DateTime              util.Time         `json:"dateTime" bson:"dateTime"`
	Meta                  *models.Meta      //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
