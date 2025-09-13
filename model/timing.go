package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// UnitOfTime ...
type UnitOfTime string

const (
	//SECOND ..
	SECOND UnitOfTime = "SECOND"

	//MINUTE ..
	MINUTE UnitOfTime = "MINUTE"

	//HOUR ..
	HOUR UnitOfTime = "HOUR"

	//DAY ..
	DAY UnitOfTime = "DAY"

	//WEEK ..
	WEEK UnitOfTime = "WEEK"

	//MONTH ..
	MONTH UnitOfTime = "MONTH"

	//YEAR ..
	YEAR UnitOfTime = "YEAR"
)

// DaysOfWeek ...
type DaysOfWeek string

const (
	//MONDAY ..
	MONDAY DaysOfWeek = "MONDAY"

	//TUESDAY ..
	TUESDAY DaysOfWeek = "TUESDAY"

	//WEDNESDAY ..
	WEDNESDAY DaysOfWeek = "WEDNESDAY"

	//THURSDAY ..
	THURSDAY DaysOfWeek = "THURSDAY"

	//FRIDAY ..
	FRIDAY DaysOfWeek = "FRIDAY"

	//SATURDAY ..
	SATURDAY DaysOfWeek = "SATURDAY"

	//SUNDAY ..
	SUNDAY DaysOfWeek = "SUNDAY"
)

// EventTiming ...
type EventTiming string

const (
	//EVENT_HS ..
	EVENT_HS EventTiming = "EVENT_HS"

	//EVENT_WAKE ..
	EVENT_WAKE EventTiming = "EVENT_WAKE"

	//EVENT_C ..
	EVENT_C EventTiming = "EVENT_C"

	//EVENT_CM ..
	EVENT_CM EventTiming = "EVENT_CM"

	//EVENT_CD ..
	EVENT_CD EventTiming = "EVENT_CD"

	//EVENT_CV ..
	EVENT_CV EventTiming = "EVENT_CV"

	//EVENT_AC ..
	EVENT_AC EventTiming = "EVENT_AC"

	//EVENT_ACM ..
	EVENT_ACM EventTiming = "EVENT_ACM"

	//EVENT_ACD ..
	EVENT_ACD EventTiming = "EVENT_ACD"

	//EVENT_ACV ..
	EVENT_ACV EventTiming = "EVENT_ACV"

	//EVENT_PC ..
	EVENT_PC EventTiming = "EVENT_PC"

	//EVENT_PCM ..
	EVENT_PCM EventTiming = "EVENT_ACV"

	//EVENT_PCD ..
	EVENT_PCD EventTiming = "EVENT_PCD"

	//EVENT_PCV ..
	EVENT_PCV EventTiming = "EVENT_PCV"
)

// TimingAbbreviation ...
type TimingAbbreviation string

const (
	//BID ..
	BID TimingAbbreviation = "BID"

	//TID ..
	TID TimingAbbreviation = "TID"

	//QID ..
	QID TimingAbbreviation = "QID"

	//AM ..
	AM TimingAbbreviation = "AM"

	//QD ..
	QD TimingAbbreviation = "QD"

	//QOD ..
	QOD TimingAbbreviation = "QOD"

	//Q4H ..
	Q4H TimingAbbreviation = "Q4H"

	//Q6H ..
	Q6H TimingAbbreviation = "Q6H"
)

// FrequenceBoundsType ...
type FrequenceBoundsType string

const (
	//FREQUENCY_BOUNDS_DURATION ..
	FREQUENCY_BOUNDS_DURATION FrequenceBoundsType = "FREQUENCY_BOUNDS_DURATION"

	//FREQUENCY_BOUNDS_RANGE ..
	FREQUENCY_BOUNDS_RANGE FrequenceBoundsType = "FREQUENCY_BOUNDS_RANGE"

	//FREQUENCY_BOUNDS_PERIOD ..
	FREQUENCY_BOUNDS_PERIOD FrequenceBoundsType = "FREQUENCY_BOUNDS_PERIOD"
)

// TimingInput ...
type TimingInput struct {
	Event          *util.Time           `json:"event"`
	BoundsType     *FrequenceBoundsType `json:"boundsType"`
	BoundsDuration *int32               `json:"boundsDuration"`
	BoundsRange    *RangeInput          `json:"boundsRange"`
	BoundsPeriod   *PeriodInput         `json:"boundsPeriod"`
	Count          *int32               `json:"count"`
	CountMax       *int32               `json:"countMax"`
	Duration       *float64             `json:"duration"`
	DurationMax    *float64             `json:"durationMax"`
	DurationUnit   *UnitOfTime          `json:"durationUnit"`
	Frequency      *int32               `json:"frequency"`
	FrequencyMax   *int32               `json:"frequencyMax"`
	Period         *float64             `json:"period"`
	PeriodMax      *float64             `json:"periodMax"`
	PeriodUnit     *UnitOfTime          `json:"periodUnit"`
	DayOfWeek      *DaysOfWeek          `json:"dayOfWeek"`
	Time           *util.Time           `json:"time" bson:"time"`
	When           *EventTiming         `json:"when" bson:"when"`
	Offset         *int32               `json:"offset" bson:"offset"`
	Code           *TimingAbbreviation  `json:"code" bson:"code"`
	Meta           *models.Meta         //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// Timing ...
type Timing struct {
	Id             string               `json:"id" bson:"_id"`
	Event          *util.Time           `json:"event" bson:"event"`
	BoundsType     *FrequenceBoundsType `json:"boundsType" bson:"boundsType"`
	BoundsDuration *int32               `json:"boundsDuration" bson:"boundsDuration"`
	BoundsRange    *Range               `json:"boundsRange" bson:"boundsRange"`
	BoundsPeriod   *Period              `json:"boundsPeriod" bson:"boundsPeriod"`
	Count          *int32               `json:"count" bson:"count"`
	CountMax       *int32               `json:"countMax" bson:"countMax"`
	Duration       *float64             `json:"duration" bson:"duration"`
	DurationMax    *float64             `json:"durationMax" bson:"durationMax"`
	DurationUnit   *UnitOfTime          `json:"durationUnit" bson:"durationUnit"`
	Frequency      *int32               `json:"frequency" bson:"frequency"`
	FrequencyMax   *int32               `json:"frequencyMax" bson:"frequencyMax"`
	Period         *float64             `json:"period" bson:"period"`
	PeriodMax      *float64             `json:"periodMax" bson:"periodMax"`
	PeriodUnit     *UnitOfTime          `json:"periodUnit" bson:"periodUnit"`
	DayOfWeek      *DaysOfWeek          `json:"dayOfWeek" bson:"dayOfWeek"`
	Time           *util.Time           `json:"time" bson:"time"`
	When           *EventTiming         `json:"when" bson:"when"`
	Offset         *int32               `json:"offset" bson:"offset"`
	Code           *TimingAbbreviation  `json:"code" bson:"code"`
	Meta           *models.Meta         //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
