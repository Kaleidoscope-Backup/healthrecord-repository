package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// RelationshipType ...
type RelationshipType string

const (
	//FAMILY_MEMBER ..
	FAMILY_MEMBER RelationshipType = "FAMILY_MEMBER"

	//CARE_TEAM ..
	CARE_TEAM RelationshipType = "CARE_TEAM"

	//CUSTOMER ..
	CUSTOMER RelationshipType = "CUSTOMER"

	//STAFF ..
	STAFF RelationshipType = "STAFF"

	//TEST_ADMINISTRATOR ..
	TEST_ADMINISTRATOR RelationshipType = "TEST_ADMINISTRATOR"

	//PARTNER ..
	PARTNER RelationshipType = "PARTNER"

	//ACCOUNT_MANAGER ..
	ACCOUNT_MANAGER RelationshipType = "ACCOUNT_MANAGER"
)

// RelationshipCreate ....
type RelationshipCreate struct {
	Active         bool                `json:"active"`
	From           ReferenceActorInput `json:"from"`
	To             ReferenceActorInput `json:"to"`
	Label          string              `json:"label"`
	Type           *RelationshipType   `json:"type"`
	Code           *string             `json:"code"`
	Consent        *string             `json:"consent"`
	Period         *PeriodInput        `json:"period"`
	AdditionalData *[]AttributeInput   `json:"additionalData"`
}

// Relationship ....
type Relationship struct {
	Id             string            `json:"id" bson:"_id"`
	Active         bool              `json:"active" bson:"active"`
	From           ReferenceActor    `json:"from" bson:"from"`
	To             ReferenceActor    `json:"to" bson:"to"`
	Label          string            `json:"label" bson:"label"`
	Type           *RelationshipType `json:"type" bson:"type"`
	Code           *ClinicalCode     `json:"code" bson:"code"`
	Consent        *string           `json:"consent" bson:"consent"`
	Period         *Period           `json:"period" bson:"period"`
	AdditionalData *[]Attribute      `json:"additionalData" bson:"additionalData"`
	Meta           *models.Meta      //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

type RelationshipQueryParam struct {
	FromID   *string
	FromType *ActorType
	ToID     *string
	ToType   *ActorType
	RelType  *RelationshipType
	Label    *string
}
