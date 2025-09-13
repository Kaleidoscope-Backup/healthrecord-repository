package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// SocialHistoryObservationType ...
type SocialHistoryObservationType string

const (
	//SMOKING ..
	SMOKING SocialHistoryObservationType = "SMOKING"

	//ALCOHOL ..
	ALCOHOL SocialHistoryObservationType = "ALCOHOL"

	//RECREATIONAL_DRUG ..
	RECREATIONAL_DRUG SocialHistoryObservationType = "RECREATIONAL_DRUG"

	//INTERNET_SURFING ..
	INTERNET_SURFING SocialHistoryObservationType = "INTERNET_SURFING"

	//SEX_LIFE ...
	SEX_LIFE SocialHistoryObservationType = "SEX_LIFE"

	//OTHER ..
	OTHER SocialHistoryObservationType = "OTHER"
)

// SocialHistoryObservationStatus ...
type SocialHistoryObservationStatus string

const (
	//NEVER ..
	NEVER SocialHistoryObservationStatus = "NEVER"

	//INFREQUENT ..
	INFREQUENT SocialHistoryObservationStatus = "INFREQUENT"

	//LIGHT ..
	LIGHT SocialHistoryObservationStatus = "LIGHT"

	//MODERATE ..
	MODERATE SocialHistoryObservationStatus = "MODERATE"

	//HEAVY ..
	HEAVY SocialHistoryObservationStatus = "HEAVY"

	//VERY_HEAVY ..
	VERY_HEAVY SocialHistoryObservationStatus = "VERY_HEAVY"

	//QUIT ..
	QUIT SocialHistoryObservationStatus = "QUIT"

	//UNKNOWN ..
	UNKNOWN SocialHistoryObservationStatus = "UNKNOWN"
)

func (t SocialHistoryObservationStatus) toDescription() string {
	switch t {
	case NEVER:
		return "Never taken"
	case INFREQUENT:
		return "Infrequent use"
	case HEAVY:
		return "Heavy frequent use"
	case QUIT:
		return "Used to take in past"
	case UNKNOWN:
		return "Status not available"

	}

	return ""
}

// SocialHistoryObservationRecordQueryParam ...
type SocialHistoryObservationRecordQueryParam struct {
	ConsumerID string                          `json:"consumerID"`
	Type       *SocialHistoryObservationType   `json:"type"`
	Status     *SocialHistoryObservationStatus `json:"status"`
}

// SocialHistoryObservationRecordCreate ...
type SocialHistoryObservationRecordCreate struct {
	HealthRecordCreate
	Type         SocialHistoryObservationType   `json:"type"`
	Status       SocialHistoryObservationStatus `json:"status"`
	Duration     *int32                         `json:"duration"`
	DurationUnit *string                        `json:"durationUnit"`
	Value        *ValueInput                    `json:"value" bson:"value"`
	Code         *CodableConceptInput           `json:"code"`
	Start        *util.Time                     `json:"start"`
	End          *util.Time                     `json:"end"`
}

// SocialHistoryObservationRecord ...
type SocialHistoryObservationRecord struct {
	HealthRecord
	Id           string                         `json:"id" bson:"_id"`
	Type         SocialHistoryObservationType   `json:"addictionType" bson:"type"`
	Status       SocialHistoryObservationStatus `json:"addictionStatus" bson:"status"`
	Duration     *int32                         `json:"duration" bson:"duration"`
	DurationUnit *string                        `json:"durationUnit" bson:"durationUnit"`
	Value        *Value                         `json:"value" bson:"value"`
	Code         *CodableConcept                `json:"code" bson:"code"`
	Start        *util.Time                     `json:"start" bson:"start"`
	End          *util.Time                     `json:"end" bson:"end"`
	Meta         *models.Meta                   //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
