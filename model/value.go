package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//ValueType ...
type ValueType string

const (
	//QUANTITY ..
	QUANTITY ValueType = "QUANTITY"

	//DECIMAL ..
	DECIMAL ValueType = "DECIMAL"

	//BOOLEAN ..
	BOOLEAN ValueType = "BOOLEAN"

	//RANGE ..
	RANGE ValueType = "RANGE"

	//RATIO ..
	RATIO ValueType = "RATIO"

	//TEXT ..
	TEXT ValueType = "TEXT"

	//DATE_TIME ..
	DATE_TIME ValueType = "DATE_TIME"

	//PERIOD ..
	PERIOD ValueType = "PERIOD"

	//RATING ..
	RATING ValueType = "RATING"

	//REFERENCE_ENTITY ..
	REFERENCE_ENTITY ValueType = "REFERENCE_ENTITY"
)

//RatingInput ...
type RatingInput struct {
	Min         int32 `json:"min"`
	Max         int32 `json:"max"`
	RatingValue int32 `json:"ratingValue"`
}

//Rating ...
type Rating struct {
	Id          string       `json:"id" bson:"_id"`
	Min         int32        `json:"min" bson:"min"`
	Max         int32        `json:"max" bson:"max"`
	RatingValue int32        `json:"ratingValue" bson:"ratingValue"`
	Meta        *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//RangeInput ...
type RangeInput struct {
	Min int32 `json:"min"`
	Max int32 `json:"max"`
}

//Range ...
type Range struct {
	Id   string       `json:"id" bson:"_id"`
	Min  int32        `json:"min" bson:"min"`
	Max  int32        `json:"max" bson:"max"`
	Meta *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//PeriodInput ...
type PeriodInput struct {
	Start util.Time `json:"start" bson:"start"`
	End   util.Time `json:"end" bson:"end"`
}

//Period ...
type Period struct {
	Id    string    `json:"id" bson:"_id"`
	Start util.Time `json:"start" bson:"start"`
	End   util.Time `json:"end" bson:"end"`
}

//RatioInput ...
type RatioInput struct {
	Numerator   int32 `json:"numerator"`
	Denominator int32 `json:"denominator"`
}

//Ratio ...
type Ratio struct {
	Id          string       `json:"id" bson:"_id"`
	Numerator   int32        `json:"numerator" bson:"numerator"`
	Denominator int32        `json:"denominator" bson:"denominator"`
	Meta        *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//ValueInput ...
type ValueInput struct {
	ValueType            ValueType             `json:"valueType"`
	ValueQuantity        *int32                `json:"valueQuantity"`
	ValueDecimal         *float64              `json:"valueDecimal"`
	ValueBoolean         *bool                 `json:"valueBoolean"`
	ValueRange           *RangeInput           `json:"valueRange"`
	ValueRatio           *RatioInput           `json:"valueRatio"`
	ValueText            *string               `json:"valueText"`
	ValueDate            *util.Time            `json:"valueDate" bson:"valueDate"`
	ValuePeriod          *PeriodInput          `json:"valuePeriod"`
	ValueRating          *RatingInput          `json:"valueRating"`
	ValueReferenceEntity *ReferenceEntityInput `json:"valueReferenceEntity"`
	Unit                 *string               `json:"unit"`
}

//Value ...
type Value struct {
	Id                   string           `json:"id" bson:"_id"`
	ValueType            ValueType        `json:"valueType" bson:"valueType"`
	ValueQuantity        *int32           `json:"valueQuantity" bson:"valueQuantity"`
	ValueDecimal         *float64         `json:"valueDecimal" bson:"valueDecimal"`
	ValueRange           *Range           `json:"valueRange" bson:"valueRange"`
	ValueRatio           *Ratio           `json:"valueRatio" bson:"valueRatio"`
	ValueBoolean         *bool            `json:"valueBoolean" bson:"valueBoolean"`
	ValueText            *string          `json:"valueText" bson:"valueText"`
	ValueDate            *util.Time       `json:"valueDate" bson:"valueDate"`
	ValuePeriod          *Period          `json:"valuePeriod" bson:"valuePeriod"`
	ValueRating          *Rating          `json:"valueRating" bson:"valueRating"`
	ValueReferenceEntity *ReferenceEntity `json:"valueReferenceEntity" bson:"valueReferenceEntity"`
	Unit                 *string          `json:"unit" bson:"unit"`
	Meta                 *models.Meta     //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
