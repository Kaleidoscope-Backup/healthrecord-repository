package model

import (
	"github.com/karte/mongo-lib/models"
)

// ObservationStatus ...
type ObservationStatus string

const (
	//OBSERVATION_REGISTERED ..
	OBSERVATION_REGISTERED ObservationStatus = "OBSERVATION_REGISTERED"

	//OBSERVATION_PRELIMINARY ..
	OBSERVATION_PRELIMINARY ObservationStatus = "OBSERVATION_PRELIMINARY"

	//OBSERVATION_FINAL ..
	OBSERVATION_FINAL ObservationStatus = "OBSERVATION_FINAL"

	//OBSERVATION_AMENDED ..
	OBSERVATION_AMENDED ObservationStatus = "OBSERVATION_AMENDED"

	//OBSERVATION_CORRECTED ..
	OBSERVATION_CORRECTED ObservationStatus = "OBSERVATION_CORRECTED"

	//OBSERVATION_CANCELLED ..
	OBSERVATION_CANCELLED ObservationStatus = "OBSERVATION_CANCELLED"

	//OBSERVATION_ENTERED_IN_ERROR ..
	OBSERVATION_ENTERED_IN_ERROR ObservationStatus = "OBSERVATION_ENTERED_IN_ERROR"

	//OBSERVATION_UNKNOWN ..
	OBSERVATION_UNKNOWN ObservationStatus = "OBSERVATION_UNKNOWN"
)

// ObservationCategory ...
type ObservationCategory string

const (
	//OBSERVATION_CATEGORY_SOCIALHISTORY ..
	OBSERVATION_CATEGORY_SOCIALHISTORY ObservationCategory = "OBSERVATION_CATEGORY_SOCIALHISTORY"

	//OBSERVATION_CATEGORY_VITALSIGNS ..
	OBSERVATION_CATEGORY_VITALSIGNS ObservationCategory = "OBSERVATION_CATEGORY_VITALSIGNS"

	//OBSERVATION_CATEGORY_IMAGING ..
	OBSERVATION_CATEGORY_IMAGING ObservationCategory = "OBSERVATION_CATEGORY_IMAGING"

	//OBSERVATION_CATEGORY_PROCEDURE ..
	OBSERVATION_CATEGORY_PROCEDURE ObservationCategory = "OBSERVATION_CATEGORY_PROCEDURE"

	//OBSERVATION_CATEGORY_SURVEY ..
	OBSERVATION_CATEGORY_SURVEY ObservationCategory = "OBSERVATION_CATEGORY_SURVEY"

	//OBSERVATION_CATEGORY_EXAM ..
	OBSERVATION_CATEGORY_EXAM ObservationCategory = "OBSERVATION_CATEGORY_EXAM"

	//OBSERVATION_CATEGORY_THERAPY ..
	OBSERVATION_CATEGORY_THERAPY ObservationCategory = "OBSERVATION_CATEGORY_THERAPY"
)

// ObservationRecordsCreate ...
type ObservationRecordsCreate struct {
	Observations *[]ObservationRecordCreate `json:"observations"`
}

// ObservationRecordCreate ...
type ObservationRecordCreate struct {
	HealthRecordCreate
	Status               ObservationStatus      `json:"status"`
	Category             ObservationCategory    `json:"category"`
	CategoryCode         *CodableConceptInput   `json:"categoryCode"`
	Code                 *CodableConceptInput   `json:"code"`
	Performer            *ReferenceActorInput   `json:"performer"`
	Effective            *PeriodInput           `json:"effective"`
	Value                ValueInput             `json:"value"`
	DataAbsentReason     *string                `json:"dataAbsentReason"`
	DataAbsentReasonCode *CodableConceptInput   `json:"dataAbsentReasonCode"`
	Interpretation       *string                `json:"interpretationCode"`
	InterpretationCode   *CodableConceptInput   `json:"interpretation"`
	Comment              *string                `json:"comment"`
	BodySite             *string                `json:"bodySite"`
	BodySiteCode         *CodableConceptInput   `json:"bodySiteCode"`
	Method               *string                `json:"method"`
	MethodCode           *CodableConceptInput   `json:"methodCode"`
	Device               *ReferenceEntityInput  `json:"device"`
	ReferenceRange       *[]ReferenceRangeInput `json:"referenceRange"`
}

// ObservationRecord ...
type ObservationRecord struct {
	HealthRecord
	Id                   string              `json:"id" bson:"_id"`
	Status               ObservationStatus   `json:"status" bson:"status"`
	Category             ObservationCategory `json:"category" bson:"category"`
	CategoryCode         *CodableConcept     `json:"categoryCode" bson:"categoryCode"`
	Code                 *CodableConcept     `json:"code" bson:"code"`
	Performer            *ReferenceActor     `json:"performer" bson:"performer"`
	Effective            *Period             `json:"effective" bson:"effective"`
	Value                Value               `json:"value" bson:"value"`
	DataAbsentReason     *string             `json:"dataAbsentReason" bson:"dataAbsentReason"`
	DataAbsentReasonCode *CodableConcept     `json:"dataAbsentReasonCode" bson:"dataAbsentReasonCode"`
	Interpretation       *string             `json:"interpretation" bson:"interpretation"`
	InterpretationCode   *CodableConcept     `json:"interpretationCode" bson:"interpretationCode"`
	Comment              *string             `json:"comment" bson:"comment"`
	BodySite             *string             `json:"bodySite" bson:"bodySite"`
	BodySiteCode         *CodableConcept     `json:"bodySiteCode" bson:"bodySiteCode"`
	Method               *string             `json:"method" bson:"method"`
	MethodCode           *CodableConcept     `json:"methodCode" bson:"methodCode"`
	Device               *ReferenceEntity    `json:"device" bson:"device"`
	ReferenceRange       *[]ReferenceRange   `json:"referenceRange" bson:"referenceRange"`
	Meta                 *models.Meta        //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
