package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// EncounterRecord ...
type EncounterRecord struct {
	HealthRecord
	Id            string            `json:"id" bson:"_id"`
	Reasons       *[]Reason         `json:"reasons" bson:"reasons"`
	Diagnosis     *[]Diagnosis      `json:"diagnosis" bson:"diagnosis"`
	Prescriptions *[]Medication     `json:"prescriptions" bson:"prescriptions"`
	Orders        *[]EncounterOrder `json:"orders" bson:"orders"`
	AttendedBy    *Practitioner     `json:"attendedBy" bson:"attendedBy"`
	Meta          *models.Meta      //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// EncounterOrder ...
type EncounterOrder struct {
	Name           string        `json:"name" bson:"name"`
	ProcedureCode  ProcedureCode `json:"procedureCode" bson:"procedureCode"`
	Code           *ClinicalCode `json:"clinicalCode" bson:"clinicalCode"`
	ExpectedDate   *util.Time    `json:"expectedDate" bson:"expectedDate"`
	ExpirationDate *util.Time    `json:"expirationDate" bson:"expirationDate"`
	Type           *string       `json:"type" bson:"type"`
}

// EncounterRecordCreate ...
type EncounterRecordCreate struct {
	HealthRecordCreate
	Reasons              *[]ReasonInput          `json:"reasons,omitempty"`
	Diagnosis            *[]DiagnosisInput       `json:"diagnosis,omitempty"`
	Prescriptions        *[]MedicationCreate     `json:"prescriptions,omitempty"`
	Orders               *[]EncounterOrderCreate `json:"orders,omitempty"`
	AttendedByID         *string                 `json:"attendedBy,omitempty"`
	SourceRecordIDSystem *string                 `json:"sourceRecordIDSystem,omitempty"`
	SourceRecordIDValue  *string                 `json:"sourceRecordIDValue,omitempty"`
}

// EncounterOrderCreate ...
type EncounterOrderCreate struct {
	Name           string        `json:"name"`
	ProcedureCode  ProcedureCode `json:"procedureCode"`
	ExpectedDate   *util.Time    `json:"expectedDate,omitempty"`
	ExpirationDate *util.Time    `json:"expirationDate,omitempty"`
	Type           *string       `json:"type,omitempty"`
}
