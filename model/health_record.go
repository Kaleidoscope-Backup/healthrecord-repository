package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// HealthRecordType ...
type HealthRecordType string

const (
	SLEEP                                HealthRecordType = "SLEEP"
	GOAL                                 HealthRecordType = "GOAL"
	MEAL                                 HealthRecordType = "MEAL"
	ADVERSE_EVENT                        HealthRecordType = "ADVERSE_EVENT"
	ACTIVITY                             HealthRecordType = "ACTIVITY"
	MEDICATION                           HealthRecordType = "MEDICATION"
	CONDITION                            HealthRecordType = "CONDITION"
	ENCOUNTER                            HealthRecordType = "ENCOUNTER"
	OBSERVATION                          HealthRecordType = "OBSERVATION"
	HEART_RATE                           HealthRecordType = "HEART_RATE"
	BLOOD_GLUCOSE                        HealthRecordType = "BLOOD_GLUCOSE"
	BLOOD_PRESSURE                       HealthRecordType = "BLOOD_PRESSURE"
	WEIGHT                               HealthRecordType = "WEIGHT"
	HEIGHT                               HealthRecordType = "HEIGHT"
	IMMUNIZATION                         HealthRecordType = "IMMUNIZATION"
	ALLERGY                              HealthRecordType = "ALLERGY"
	SPECIMEN                             HealthRecordType = "SPECIMEN"
	DIAGNOSTIC_REPORT                    HealthRecordType = "DIAGNOSTIC_REPORT"
	FAMILY_HISTORY                       HealthRecordType = "FAMILY_HISTORY"
	ADDICTION                            HealthRecordType = "ADDICTION"
	PROCEDURE                            HealthRecordType = "PROCEDURE"
	PERSONAL_CHARACTERISTICS_OBSERVATION HealthRecordType = "PERSONAL_CHARACTERISTICS_OBSERVATION"
	CLINICAL_ASSESMENT_OBSERVATION       HealthRecordType = "CLINICAL_ASSESMENT_OBSERVATION"
	IMAGING_RESULT_OBSERVATION           HealthRecordType = "IMAGING_RESULT_OBSERVATION"
	LAB_RESULT_OBSERVATION               HealthRecordType = "LAB_RESULT_OBSERVATION"
	VITAL_OBSERVATION                    HealthRecordType = "VITAL_OBSERVATION"
	SOCIAL_HISTORY_OBSERVATION_RECORD    HealthRecordType = "SOCIAL_HISTORY_OBSERVATION_RECORD"
	APPOINTMENT                          HealthRecordType = "APPOINTMENT"
	NUTRITION_ORDER                      HealthRecordType = "NUTRITION_ORDER"
	MOLECULAR_SEQUENCE                   HealthRecordType = "MOLECULAR_SEQUENCE"
)

// HealthRecordTransactionType ...
type HealthRecordTransactionType string

const (
	INSERT HealthRecordTransactionType = "INSERT"
	CREATE HealthRecordTransactionType = "CREATE"
	UPDATE HealthRecordTransactionType = "UPDATE"
	DELETE HealthRecordTransactionType = "DELETE"
)

// ReferenceHealthRecordInput ...
type ReferenceHealthRecordInput struct {
	Type         HealthRecordType `bson:"type,omitempty"`
	ReferencedID string           `bson:"referenceid,omitempty"`
}

// ReferenceHealthRecord ...
type ReferenceHealthRecord struct {
	Id           string           `json:"id" bson:"_id"`
	Type         HealthRecordType `bson:"type,omitempty" json:"type,omitempty"`
	ReferencedID string           `bson:"referenceid,omitempty" json:"referenceid,omitempty"`
	Meta         *models.Meta     //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// HealthRecordCreate ...
type HealthRecordCreate struct {
	ConsumerID     string                        `json:"consumerID"`
	Name           string                        `json:"name"`
	Description    *string                       `json:"description,omitempty"`
	Occurred       util.Time                     `json:"occurred"`
	Source         string                        `json:"source"`
	Organization   *string                       `json:"organization,omitempty"`
	SourceRecordID *SourceRecordIDInput          `json:"sourceRecordID,omitempty"`
	References     *[]ReferenceHealthRecordInput `json:"references"`
	Location       *GeoLocationInput             `json:"location"`
}

// HealthRecord ...
type HealthRecord struct {
	Id              string                      `json:"id" bson:"_id,omitempty"`
	ConsumerID      string                      `json:"consumerID" bson:"consumerID"`
	RecordType      HealthRecordType            `json:"recordType" bson:"recordType"`
	TransactionType HealthRecordTransactionType `json:"transactionType" bson:"transactionType"`
	Name            string                      `json:"name" bson:"name"`
	Description     *string                     `json:"description,omitempty" bson:"description,omitempty"`
	Occurred        util.Time                   `json:"occurred" bson:"occurred"`
	Created         util.Time                   `json:"created" bson:"created"`
	CreatedBy       *string                     `json:"createdBy" bson:"createdBy"`
	Source          string                      `json:"source" bson:"source"`
	PreviousRecord  *string                     `json:"previousRecord,omitempty" bson:"previousRecord,omitempty"`
	Organization    *string                     `json:"organization,omitempty" bson:"organization,omitempty"`
	SourceRecordID  *SourceRecordID             `json:"sourceRecordID" bson:"sourceRecordID,omitempty"`
	References      *[]ReferenceHealthRecord    `json:"references" bson:"references"`
	Location        *GeoLocation                `json:"location" bson:"location"`
	Meta            *models.Meta                //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// HealthRecords ...
type HealthRecords struct {
	Count   int32          `json:"count" bson:"count"`
	Records []HealthRecord `json:"records" bson:"records"`
	Meta    *models.Meta   //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// Diagnosis ...
type Diagnosis struct {
	Name string        `json:"name" bson:"name"`
	Code *ClinicalCode `json:"clinicalCode" bson:"clinicalCode"`
}

// DiagnosisInput ...
type DiagnosisInput struct {
	Name string `json:"name"`
}

// Reason ...
type Reason struct {
	Name string        `json:"name" bson:"name"`
	Code *ClinicalCode `json:"clinicalCode" bson:"clinicalCode"`
}

// ReasonInput ...
type ReasonInput struct {
	Name string `json:"name"`
}
