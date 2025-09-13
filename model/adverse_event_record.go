package model

import "gitlab.com/karte/mongo-lib/models"

//AdverseEventCategory ...
type AdverseEventCategory string

const (
	//AE ..
	AE AdverseEventCategory = "AE"
	//PAE ..
	PAE AdverseEventCategory = "PAE"
	//AE_OTHER ..
	AE_OTHER AdverseEventCategory = "AE_OTHER"
)

//AdverseEventRecordCreate ...
type AdverseEventRecordCreate struct {
	HealthRecordCreate
	Category      AdverseEventCategory `json:"category"`
	CategoryCode  *CodableConceptInput `json:"categoryCode"`
	EventType     string               `json:"eventType"`
	EventTypeCode *CodableConceptInput `json:"typeCode"`
	Location      *GeoLocationInput    `json:"location"`
	Seriousness   *Severity            `json:"seriousness"`
	Outcome       *string              `json:"outcome"`
	OutcomeCode   *CodableConceptInput `json:"outcomeCode"`
	Recorder      *ReferenceActorInput `json:"recorder"`
}

//AdverseEventRecord ...
type AdverseEventRecord struct {
	HealthRecord
	Id            string               `json:"id" bson:"_id"`
	Category      AdverseEventCategory `json:"category" bson:"category"`
	CategoryCode  *CodableConcept      `json:"categoryCode" bson:"categoryCode"`
	EventType     string               `json:"eventType" bson:"eventType"`
	EventTypeCode *CodableConcept      `json:"typeCode" bson:"typeCode"`
	Location      *GeoLocation         `json:"location" bson:"location"`
	Seriousness   *Severity            `json:"seriousness" bson:"seriousness"`
	Outcome       *string              `json:"outcome" bson:"outcome"`
	OutcomeCode   *CodableConcept      `json:"outcomeCode" bson:"outcomeCode"`
	Recorder      *ReferenceActor      `json:"recorder" bson:"recorder"`
	Meta          *models.Meta         //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
