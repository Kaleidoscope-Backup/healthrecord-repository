package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//GoalCategory ...
type GoalCategory string

const (
	//DIETARY_GOAL ..
	DIETARY_GOAL GoalCategory = "DIETARY_GOAL"
	//SAFETY_GOAL ..
	SAFETY_GOAL GoalCategory = "SAFETY_GOAL"
	//BEHAVIORAL_GOAL ..
	BEHAVIORAL_GOAL GoalCategory = "BEHAVIORAL_GOAL"
	//NURSING_GOAL ..
	NURSING_GOAL GoalCategory = "NURSING_GOAL"
	//PHYSIOTHERAPY_GOAL ..
	PHYSIOTHERAPY_GOAL GoalCategory = "PHYSIOTHERAPY_GOAL"
	//OTHER_GOAL ..
	OTHER_GOAL GoalCategory = "OTHER_GOAL"
)

//GoalRecordCreate ...
type GoalRecordCreate struct {
	HealthRecordCreate
	Category     GoalCategory                  `json:"goalCategory"`
	CategoryCode *CodableConceptInput          `json:"categoryCode"`
	Priority     *Priority                     `json:"priority"`
	Start        util.Time                     `json:"start"`
	DueDate      *util.Time                    `json:"dueDate"`
	DueDuration  *int32                        `json:"dueDuration"`
	Measure      string                        `json:"measure"`
	MeasureCode  *CodableConceptInput          `json:"measureCode"`
	Target       ValueInput                    `json:"target"`
	ExpressedBy  *ReferenceActorInput          `json:"expressedBy"`
	Outcomes     *[]ReferenceHealthRecordInput `json:"outcomes"`
	Note         *string                       `json:"note"`
}

//GoalRecord ...
type GoalRecord struct {
	HealthRecord
	Id           string                   `json:"id" bson:"_id"`
	Category     GoalCategory             `json:"goalCategory" bson:"goalCategory"`
	CategoryCode *CodableConcept          `json:"categoryCode" bson:"categoryCode"`
	Priority     *Priority                `json:"priority" bson:"priority"`
	Start        util.Time                `json:"start" bson:"start"`
	DueDate      *util.Time               `json:"dueDate" bson:"dueDate"`
	DueDuration  *int32                   `json:"dueDuration" bson:"dueDuration"`
	Measure      string                   `json:"measure" bson:"measure"`
	MeasureCode  *CodableConcept          `json:"measureCode" bson:"measureCode"`
	Target       Value                    `json:"target" bson:"target"`
	ExpressedBy  *ReferenceActor          `json:"expressedBy" bson:"expressedBy"`
	Outcomes     *[]ReferenceHealthRecord `json:"outcomes" bson:"outcomes"`
	Note         *string                  `json:"note" bson:"note"`
	Meta         *models.Meta             //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
