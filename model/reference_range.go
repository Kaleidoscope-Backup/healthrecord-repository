package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// ReferenceRangeType ...
type ReferenceRangeType string

const (
	//NORMAL ..
	NORMAL ReferenceRangeType = "NORMAL"

	//RECOMMENDED ..
	RECOMMENDED ReferenceRangeType = "RECOMMENDED"

	//HIGH ..
	HIGH ReferenceRangeType = "HIGH"

	//VERY_HIGH ..
	VERY_HIGH ReferenceRangeType = "VERY_HIGH"

	//LOW ..
	LOW ReferenceRangeType = "LOW"

	//VERY_LOW ..
	VERY_LOW ReferenceRangeType = "VERY_LOW"

	//NOT_RECOMMENDED ..
	NOT_RECOMMENDED ReferenceRangeType = "NOT_RECOMMENDED"
)

// ReferenceRangeInput ...
type ReferenceRangeInput struct {
	Range         *RangeInput          `json:"range"`
	LowerLimit    *int32               `json:"lowerLimit"`
	HigherLimit   *int32               `json:"higherLimit"`
	RangeUnit     *string              `json:"rangeUnit"`
	AgeMin        *int32               `json:"ageMin"`
	AgeMax        *int32               `json:"ageMax"`
	AgeGroup      *[]AgeGroup          `json:"ageGroup"`
	AgeUnit       *string              `json:"ageUnit"`
	Type          ReferenceRangeType   `json:"referenceRangeType"`
	AppliesTo     *[]string            `json:"appliesTo"`
	AppliesToCode *[]ClinicalCodeInput `json:"appliesToCode"`
	Meta          *models.Meta         //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// ReferenceRange ...
type ReferenceRange struct {
	Id            string             `json:"id" bson:"_id"`
	Range         *Range             `json:"range" bson:"range"`
	LowerLimit    *int32             `json:"lowerLimit" bson:"lowerLimit"`
	HigherLimit   *int32             `json:"higherLimit" bson:"higherLimit"`
	RangeUnit     *string            `json:"rangeUnit" bson:"rangeUnit"`
	AgeMin        *int32             `json:"ageMin" bson:"ageMin"`
	AgeMax        *int32             `json:"ageMax" bson:"ageMax"`
	AgeUnit       *string            `json:"ageUnit" bson:"ageUnit"`
	AgeGroup      *[]AgeGroup        `json:"ageGroup" bson:"ageGroup"`
	Type          ReferenceRangeType `json:"referenceRangeType" bson:"referenceRangeType"`
	AppliesTo     *[]string          `json:"appliesTo" bson:"appliesTo"`
	AppliesToCode *[]ClinicalCode    `json:"appliesToCode" bson:"appliesToCode"`
	Meta          *models.Meta       //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
