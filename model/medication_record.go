package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//MedicationRecordCreate ...
type MedicationRecordCreate struct {
	HealthRecordCreate
	DispensingOrganization *string             `json:"dispensingOrganization"`
	PrescribedBy           *string             `json:"prescribedBy"`
	PrescribedOn           *util.Time          `json:"prescribedOn,omitempty"`
	Expiration             *util.Time          `json:"expiration,omitempty"`
	Medications            *[]MedicationCreate `json:"medications"`
}

//MedicationRecord ...
type MedicationRecord struct {
	HealthRecord
	Id                     string        `json:"id" bson:"_id"`
	DispensingOrganization *string       `json:"dispensingOrganization" bson:"dispensingOrganization"`
	PrescribedBy           *string       `json:"prescribedBy" bson:"prescribedBy"`
	PrescribedOn           *util.Time    `json:"prescribedOn,omitempty" bson:"prescribedOn,omitempty"`
	Expiration             *util.Time    `json:"expiration,omitempty" bson:"expiration,omitempty"`
	Medications            *[]Medication `json:"medications" bson:"medications"`
	Meta                   *models.Meta  //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
