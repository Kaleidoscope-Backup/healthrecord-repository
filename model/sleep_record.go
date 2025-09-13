package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// SleepStageType ...
type SleepStageType string

const (
	//STAGE_WAKE ..
	STAGE_WAKE SleepStageType = "STAGE_WAKE"
	//STAGE_REM ..
	STAGE_REM SleepStageType = "STAGE_REM"
	//STAGE_N1 ..
	STAGE_N1 SleepStageType = "STAGE_N1"
	//STAGE_N2 ..
	STAGE_N2 SleepStageType = "STAGE_N2"
	//STAGE_N3 ..
	STAGE_N3 SleepStageType = "STAGE_N3"
	//STAGE_N4 ..
	STAGE_N4 SleepStageType = "STAGE_N4"
)

// Sleep stage ...
type SleepStage struct {
	Id              string          `json:"id" bson:"_id"`
	Type            *SleepStageType `json:"type" bson:"type"`
	Duration        *float64        `json:"duration" bson:"duration"`
	Latency         *float64        `json:"latency" bson:"latency"`
	TotalSleepTime  *float64        `json:"totalSleepTime" bson:"totalSleepTime"`
	SleepPeriodTime *float64        `json:"sleepPeriodTime" bson:"sleepPeriodTime"`
	Meta            *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// Sleep stage input ...
type SleepStageInput struct {
	Type            *SleepStageType `json:"type"`
	Duration        *float64        `json:"duration"`
	Latency         *float64        `json:"latency"`
	TotalSleepTime  *float64        `json:"totalSleepTime"`
	SleepPeriodTime *float64        `json:"sleepPeriodTime"`
}

// SleepContinuity ...
type SleepContinuity struct {
	Id              string       `json:"id" bson:"_id"`
	SourceOfArousal *string      `json:"sourceOfArousal" bson:"sourceOfArousal"`
	NermCount       *int32       `json:"nermCount" bson:"nermCount"`
	NermIndex       *float64     `json:"nermIndex" bson:"nermIndex"`
	RemCount        *int32       `json:"remCount" bson:"remCount"`
	RemIndex        *float64     `json:"remIndex" bson:"remIndex"`
	TotalCount      *int32       `json:"totalCount" bson:"totalCount"`
	TotalIndex      *float64     `json:"totalIndex" bson:"totalIndex"`
	Meta            *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// SleepContinuityInput ...
type SleepContinuityInput struct {
	SourceOfArousal *string  `json:"sourceOfArousal"`
	NermCount       *int32   `json:"nermCount"`
	NermIndex       *float64 `json:"nermIndex"`
	RemCount        *int32   `json:"remCount"`
	RemIndex        *float64 `json:"remIndex"`
	TotalCount      *int32   `json:"totalCount"`
	TotalIndex      *float64 `json:"totalIndex"`
}

// SleepRecordCreate ...
type SleepRecordCreate struct {
	HealthRecordCreate
	StartTime          util.Time               `json:"startTime"`
	EndTime            util.Time               `json:"endTime"`
	MainSleep          *bool                   `json:"mainSleep"`
	TimeUnit           *string                 `json:"timeUnit"`
	TotalRecordingTime *float64                `json:"totalRecordingTime"`
	TotalSleepTime     *float64                `json:"totalSleepTime"`
	TimeAwake          *float64                `json:"timeAwake"`
	SleepEfficiency    *float64                `json:"sleepEfficiency"`
	TimeToFallAsleep   *float64                `json:"timeToFallAsleep"`
	NumberOfAwekenings *int32                  `json:"numberOfAwekenings"`
	TimeAfterWakeup    *float64                `json:"timeAfterWakeup"`
	TimeInBed          *float64                `json:"timeInBed"`
	SleepStageSummary  *[]SleepStageInput      `json:"sleepStageSummary"`
	SleepContinuities  *[]SleepContinuityInput `json:"sleepContinuities"`
}

// SleepRecord ...
type SleepRecord struct {
	HealthRecord
	Id                 string             `json:"id" bson:"_id"`
	StartTime          util.Time          `json:"startTime" bson:"startTime"`
	EndTime            util.Time          `json:"endTime" bson:"endTime"`
	MainSleep          *bool              `json:"mainSleep" bson:"mainSleep"`
	TimeUnit           *string            `json:"timeUnit" bson:"timeUnit"`
	TotalRecordingTime *float64           `json:"totalRecordingTime" bson:"totalRecordingTime"`
	TotalSleepTime     *float64           `json:"totalSleepTime" bson:"totalSleepTime"`
	TimeAwake          *float64           `json:"timeAwake" bson:"timeAwake"`
	SleepEfficiency    *float64           `json:"sleepEfficiency" bson:"sleepEfficiency"`
	TimeToFallAsleep   *float64           `json:"timeToFallAsleep" bson:"timeToFallAsleep"`
	NumberOfAwekenings *int32             `json:"numberOfAwekenings" bson:"numberOfAwekenings"`
	TimeAfterWakeup    *float64           `json:"timeAfterWakeup" bson:"timeAfterWakeup"`
	TimeInBed          *float64           `json:"timeInBed" bson:"timeInBed"`
	SleepStageSummary  *[]SleepStage      `json:"sleepStageSummary" bson:"sleepStageSummary"`
	SleepContinuities  *[]SleepContinuity `json:"sleepContinuities" bson:"sleepContinuities"`
	Meta               *models.Meta       //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
