package model

import "gitlab.com/karte/mongo-lib/models"

//ImagingResultObservationRecordCreate ...
type ImagingResultObservationRecordCreate struct {
	HealthRecordCreate
	Code           *CodableConceptInput `json:"code"`
	Comment        *string              `json:"comment"`
	Interpretation *string              `json:"interpretation"`
	Observations   *[]AttachmentInput   `json:"observations"`
}

//ImagingResultObservationRecord ...
type ImagingResultObservationRecord struct {
	HealthRecord
	Id             string          `json:"id" bson:"_id"`
	Code           *CodableConcept `json:"code" bson:"code"`
	Comment        *string         `json:"comment" bson:"comment"`
	Interpretation *string         `json:"interpretation" bson:"interpretation"`
	Observations   *[]Attachment   `json:"observations" bson:"observations"`
	Meta           *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
