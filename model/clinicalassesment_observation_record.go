package model

import "gitlab.com/karte/mongo-lib/models"

//ClinicalAssesmentObservationCreate ...
type ClinicalAssesmentObservationCreate struct {
	Name  string               `json:"name"`
	Value string               `json:"value"`
	Score *int32               `json:"score"`
	Code  *CodableConceptInput `json:"code"`
}

//ClinicalAssesmentObservation ...
type ClinicalAssesmentObservation struct {
	Id    string          `json:"id" bson:"_id"`
	Name  string          `json:"name" bson:"name"`
	Value string          `json:"value" bson:"value"`
	Score *int32          `json:"score" bson:"score"`
	Code  *CodableConcept `json:"code" bson:"code"`
	Meta  *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//ClinicalAssesmentObservationRecordCreate ...
type ClinicalAssesmentObservationRecordCreate struct {
	HealthRecordCreate
	Code           *CodableConceptInput                  `json:"code"`
	Comment        *string                               `json:"comment"`
	Method         *string                               `json:"method"`
	MethodCode     *CodableConceptInput                  `json:"methodcode"`
	Interpretation *string                               `json:"interpretation"`
	Observations   *[]ClinicalAssesmentObservationCreate `json:"observations"`
}

//ClinicalAssesmentObservationRecord ...
type ClinicalAssesmentObservationRecord struct {
	HealthRecord
	Id             string                          `json:"id" bson:"_id"`
	Code           *CodableConcept                 `json:"code" bson:"code"`
	Comment        *string                         `json:"comment" bson:"comment"`
	Method         *string                         `json:"method" bson:"method"`
	MethodCode     *CodableConcept                 `json:"methodcode" bson:"methodcode"`
	Interpretation *string                         `json:"interpretation" bson:"interpretation"`
	Observations   *[]ClinicalAssesmentObservation `json:"observations" bson:"observations"`
	Meta           *models.Meta                    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
