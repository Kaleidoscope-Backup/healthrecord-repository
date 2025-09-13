package model

import "github.com/karte/mongo-lib/models"

// PersonalCharacteristics ...
type PersonalCharacteristics string

const (
	//EYE_COLOR ..
	EYE_COLOR PersonalCharacteristics = "EYE_COLOR"
	//HAIR_COLOR ..
	HAIR_COLOR PersonalCharacteristics = "HAIR_COLOR"
	//SKIN_COLOR ..
	SKIN_COLOR PersonalCharacteristics = "SKIN_COLOR"
	//AGE ..
	AGE PersonalCharacteristics = "AGE"
	//RACE ..
	RACE PersonalCharacteristics = "RACE"
	//GENDER ..
	GENDER PersonalCharacteristics = "GENDER"
	//BIRTHDATE ..
	BIRTHDATE PersonalCharacteristics = "BIRTHDATE"
)

// PersonalCharacteristicsObservationCreate ...
type PersonalCharacteristicsObservationCreate struct {
	Type  PersonalCharacteristics `json:"type"`
	Value string                  `json:"value"`
	Code  *CodableConceptInput    `json:"code"`
}

// PersonalCharacteristicsObservation ...
type PersonalCharacteristicsObservation struct {
	Id    string                  `json:"id" bson:"_id"`
	Type  PersonalCharacteristics `json:"type" bson:"type"`
	Value string                  `json:"value" bson:"value"`
	Code  *CodableConcept         `json:"code" bson:"code"`
	Meta  *models.Meta            //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// PersonalCharacteristicsObservationRecordCreate ...
type PersonalCharacteristicsObservationRecordCreate struct {
	HealthRecordCreate
	Observations *[]PersonalCharacteristicsObservationCreate `json:"observations"`
}

// PersonalCharacteristicsObservationRecord ...
type PersonalCharacteristicsObservationRecord struct {
	HealthRecord
	Id           string                                `json:"id" bson:"_id"`
	Observations *[]PersonalCharacteristicsObservation `json:"observations" bson:"observations"`
	Meta         *models.Meta                          //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
