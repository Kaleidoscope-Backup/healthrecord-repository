package model

import "gitlab.com/karte/mongo-lib/models"

//VitalType ...
type VitalType string

const (
	//VITAL_HEART_RATE ..
	VITAL_HEART_RATE VitalType = "VITAL_HEART_RATE"

	//VITAL_BLOOD_GLUCOSE ..
	VITAL_BLOOD_GLUCOSE VitalType = "VITAL_BLOOD_GLUCOSE"

	//VITAL_WEIGHT ..
	VITAL_WEIGHT VitalType = "VITAL_WEIGHT"

	//VITAL_HEIGHT ..
	VITAL_HEIGHT VitalType = "VITAL_HEIGHT"

	//VITAL_OXYGEN_SATURATION ..
	VITAL_OXYGEN_SATURATION VitalType = "VITAL_OXYGEN_SATURATION"

	//VITAL_TEMPERATURE ..
	VITAL_TEMPERATURE VitalType = "VITAL_TEMPERATURE"

	//VITAL_PULSERATE ..
	VITAL_PULSERATE VitalType = "VITAL_PULSERATE"

	//VITAL_RESPIRATIONRATE ..
	VITAL_RESPIRATIONRATE VitalType = "VITAL_RESPIRATIONRATE"

	//VITAL_HEAD_CIRCUMFERANCE ..
	VITAL_HEAD_CIRCUMFERANCE VitalType = "VITAL_HEAD_CIRCUMFERANCE"

	//VITAL_BODY_MASS_INDEX ..
	VITAL_BODY_MASS_INDEX VitalType = "VITAL_BODY_MASS_INDEX"

	//VITAL_SYSTOLIC_BLOODPRESSURE ..
	VITAL_SYSTOLIC_BLOODPRESSURE VitalType = "VITAL_SYSTOLIC_BLOODPRESSURE"

	//VITAL_DISTOLIC_BLOODPRESSURE ..
	VITAL_DISTOLIC_BLOODPRESSURE VitalType = "VITAL_DISTOLIC_BLOODPRESSURE"

	//VITAL_BLOODPRESSURE ..
	VITAL_BLOODPRESSURE VitalType = "VITAL_BLOODPRESSURE"

	//VITAL_CHOLESTEROL ..
	VITAL_CHOLESTEROL VitalType = "VITAL_CHOLESTEROL"
)

func (t VitalType) toDescription() string {
	switch t {
	case VITAL_HEART_RATE:
		return "Heart Rate"
	case VITAL_BLOOD_GLUCOSE:
		return "Blood Glucose"
	case VITAL_WEIGHT:
		return "Weight"
	case VITAL_HEIGHT:
		return "Height"
	}

	return ""
}

//VitalCreate ...
type VitalCreate struct {
	VitalType VitalType            `json:"vitalType"`
	Value     int32                `json:"value"`
	Unit      string               `json:"unit"`
	Code      *CodableConceptInput `json:"code"`
}

//Vital ...
type Vital struct {
	Id        string          `json:"id" bson:"_id"`
	VitalType VitalType       `json:"vitalType" bson:"vitalType"`
	Value     int32           `json:"value" bson:"value"`
	Unit      string          `json:"unit" bson:"unit"`
	Code      *CodableConcept `json:"code" bson:"code"`
	Meta      *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//VitalObservationRecordCreate ...
type VitalObservationRecordCreate struct {
	HealthRecordCreate
	Observations *[]VitalCreate `json:"observations"`
}

//VitalObservationRecord ...
type VitalObservationRecord struct {
	HealthRecord
	Id           string       `json:"id" bson:"_id"`
	Observations *[]Vital     `json:"observations" bson:"observations"`
	Meta         *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
