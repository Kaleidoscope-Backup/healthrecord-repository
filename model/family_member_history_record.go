package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// FamilyMemberHistoryCreate ...
type FamilyMemberHistoryCreate struct {
	MemberName       string               `json:"memberName"`
	Gender           *Gender              `json:"gender"`
	DateOfBirth      *util.Time           `json:"dateOfBirth"`
	Deceased         *bool                `json:"deceased"`
	Relationship     *string              `json:"relationship"`
	RelationshipCode *CodableConceptInput `json:"relationshipCode"`
	Condition        string               `json:"condition"`
	ConditionCode    *CodableConceptInput `json:"conditionCode"`
	Outcome          *string              `json:"outcome"`
	OutcomeCode      *CodableConceptInput `json:"outcomeCode"`
	Note             *string              `json:"note"`
}

// FamilyMemberHistoryRecordCreate ...
type FamilyMemberHistoryRecordCreate struct {
	HealthRecordCreate
	MemberHistory *[]FamilyMemberHistoryCreate
}

// FamilyMemberHistory ...
type FamilyMemberHistory struct {
	Id               string          `json:"id" bson:"_id"`
	MemberName       string          `json:"memberName" bson:"memberName"`
	Gender           *Gender         `json:"gender" bson:"gender"`
	DateOfBirth      *util.Time      `json:"dateOfBirth" bson:"dateOfBirth"`
	Deceased         *bool           `json:"deceased" bson:"deceased"`
	Relationship     *string         `json:"relationship" bson:"relationship"`
	RelationshipCode *CodableConcept `json:"relationshipCode" bson:"relationshipCode"`
	Condition        string          `json:"condition" bson:"condition"`
	ConditionCode    *CodableConcept `json:"conditionCode" bson:"conditionCode"`
	Outcome          *string         `json:"outcome" bson:"outcome"`
	OutcomeCode      *CodableConcept `json:"outcomeCode" bson:"outcomeCode"`
	Note             *string         `json:"note" bson:"note"`
	Meta             *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection

}

// FamilyMemberHistoryRecord ...
type FamilyMemberHistoryRecord struct {
	HealthRecord
	Id            string                 `json:"id" bson:"_id"`
	MemberHistory *[]FamilyMemberHistory `json:"memberHistory" bson:"memberHistory"`
	Meta          *models.Meta           //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection

}
