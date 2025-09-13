package model

import "github.com/karte/mongo-lib/models"

// LabResultObservationCreate ...
type LabResultObservationCreate struct {
	Name      string                 `json:"name"`
	Value     ValueInput             `json:"value"`
	Category  *string                `json:"category"`
	Code      *CodableConceptInput   `json:"code"`
	Ranges    *[]ReferenceRangeInput `json:"ranges"`
	Artifacts *[]AttachmentInput     `json:"artifacts"`
}

// LabResultObservation ...
type LabResultObservation struct {
	Id        string            `json:"id" bson:"_id"`
	Name      string            `json:"name" bson:"name"`
	Value     Value             `json:"value" bson:"value"`
	Category  *string           `json:"category" bson:"category"`
	Code      *CodableConcept   `json:"code" bson:"code"`
	Ranges    *[]ReferenceRange `json:"ranges" bson:"ranges"`
	Artifacts *[]Attachment     `json:"artifacts" bson:"artifacts"`
	Meta      *models.Meta      //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// LabResultObservationRecordQueryParam ...
type LabResultObservationRecordQueryParam struct {
	ConsumerID      string    `json:"consumerID"`
	Name            *string   `json:"name"`
	Category        *string   `json:"category"`
	ObservationName *[]string `json:"observationName"`
}

// LabResultObservationRecordCreate ...
type LabResultObservationRecordCreate struct {
	HealthRecordCreate
	Category       string                        `json:"category"`
	Code           *CodableConceptInput          `json:"code"`
	Specimen       *string                       `json:"specimen"`
	Comment        *string                       `json:"comment"`
	Method         *string                       `json:"method"`
	MethodCode     *CodableConceptInput          `json:"methodCode"`
	Interpretation *string                       `json:"interpretation"`
	Observations   *[]LabResultObservationCreate `json:"observations"`
}

// LabResultObservationRecord ...
type LabResultObservationRecord struct {
	HealthRecord
	Id             string                  `json:"id" bson:"_id"`
	Category       string                  `json:"category" bson:"category"`
	Code           *CodableConcept         `json:"code" bson:"code"`
	Specimen       *string                 `json:"specimen" bson:"specimen"`
	Comment        *string                 `json:"comment" bson:"comment"`
	Method         *string                 `json:"method" bson:"method"`
	MethodCode     *CodableConcept         `json:"methodCode" bson:"methodCode"`
	Interpretation *string                 `json:"interpretation" bson:"interpretation"`
	Observations   *[]LabResultObservation `json:"observations" bson:"observations"`
	Meta           *models.Meta            //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
