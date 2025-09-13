package model

import "github.com/Kaleidoscope-Backup/mongo-lib/models"

// ActivityType ...
type ActivityType string

const (
	//WALKING ..
	WALKING ActivityType = "WALKING"

	//WALKING_FITNESS ..
	WALKING_FITNESS ActivityType = "WALKING_FITNESS"

	//WALKING_NORDIC ..
	WALKING_NORDIC ActivityType = "WALKING_NORDIC"

	//WALKING_STROLLER ..
	WALKING_STROLLER ActivityType = "WALKING_STROLLER"

	//WALKING_TREADMILL ..
	WALKING_TREADMILL ActivityType = "WALKING_TREADMILL"

	//WHEELCHAIR ..
	WHEELCHAIR ActivityType = "WHEELCHAIR"

	//WINDSURFING ..
	WINDSURFING ActivityType = "WINDSURFING"

	//ZUMBA ..
	ZUMBA ActivityType = "ZUMBA"

	//RUNNING ..
	RUNNING ActivityType = "RUNNING"

	//CYCLING ..
	CYCLING ActivityType = "CYCLING"

	//EXCERCISE ..
	EXCERCISE ActivityType = "EXCERCISE"

	//ARCHERY ..
	ARCHERY ActivityType = "ARCHERY"

	//BOWLING ..
	BOWLING ActivityType = "BOWLING"

	//FENCING ...
	FENCING ActivityType = "FENCING"

	//GYMNASTICS ...
	GYMNASTICS ActivityType = "GYMNASTICS"

	//TRACK_AND_FIELD ...
	TRACK_AND_FIELD ActivityType = "TRACK_AND_FIELD"

	//AMERICAN_FOOTBALL ...
	AMERICAN_FOOTBALL ActivityType = "AMERICAN_FOOTBALL"

	//AUSTRELIAN_FOOTBALL ...
	AUSTRELIAN_FOOTBALL ActivityType = "AUSTRELIAN_FOOTBALL"

	//BASEBALL ...
	BASEBALL ActivityType = "BASEBALL"

	//BASKETBALL ...
	BASKETBALL ActivityType = "EXCERCISE"

	//CRICKET ...
	CRICKET ActivityType = "CRICKET"

	//HANDBALL ...
	HANDBALL ActivityType = "HANDBALL"

	//HOCKEY ...
	HOCKEY ActivityType = "HOCKEY"

	//LACROSSE ...
	LACROSSE ActivityType = "LACROSSE"

	//RUGBY ...
	RUGBY ActivityType = "RUGBY"

	//SOCCER ...
	SOCCER ActivityType = "SOCCER"

	//SOFTBALL ...
	SOFTBALL ActivityType = "SOFTBALL"

	//VOLLEYBALL ...
	VOLLEYBALL ActivityType = "VOLLEYBALL"

	//VOLLEYBALL_BEACH ...
	VOLLEYBALL_BEACH ActivityType = "VOLLEYBALL_BEACH"

	//VOLLEYBALL_INDOOR ...
	VOLLEYBALL_INDOOR ActivityType = "VOLLEYBALL_INDOOR"

	//CORE_TRAINING ...
	CORE_TRAINING ActivityType = "CORE_TRAINING"

	//ELLIPTICAL ...
	ELLIPTICAL ActivityType = "ELLIPTICAL"

	//MIXED_CARDIO ...
	MIXED_CARDIO ActivityType = "MIXED_CARDIO"

	//STAIR_CLIMBING ...
	STAIR_CLIMBING ActivityType = "STAIR_CLIMBING"

	//STAIRS ...
	STAIRS ActivityType = "STAIRS"

	//STEP_TRAINING ...
	STEP_TRAINING ActivityType = "STEP_TRAINING"

	//SLEEP_AWAKE ...
	SLEEP_AWAKE ActivityType = "SLEEP_AWAKE"

	//SLEEP_DEEP ...
	SLEEP_DEEP ActivityType = "SLEEP_DEEP"

	//SLEEP_LIGHT ...
	SLEEP_LIGHT ActivityType = "SLEEP_LIGHT"

	//SLEEP_REM ...
	SLEEP_REM ActivityType = "SLEEP_REM"

	//DANCE ...
	DANCE ActivityType = "DANCE"

	//YOGA ...
	YOGA ActivityType = "YOGA"

	//AEROBIC ...
	AEROBIC ActivityType = "ActivityType"

	//MIND_AND_BODY ...
	MIND_AND_BODY ActivityType = "MIND_AND_BODY"

	//PILATES ...
	PILATES ActivityType = "PILATES"

	//BOXING ...
	BOXING ActivityType = "BOXING"

	//KICK_BOXING ...
	KICK_BOXING ActivityType = "KICK_BOXING"

	//TAI_CHI ...
	TAI_CHI ActivityType = "TAI_CHI"

	//WRESTLING ...
	WRESTLING ActivityType = "WRESTLING"

	//BADMINTON ...
	BADMINTON ActivityType = "BADMINTON"

	//RACQUETBALL ...
	RACQUETBALL ActivityType = "RACQUETBALL"

	//SQUASH ...
	SQUASH ActivityType = "SQUASH"

	//TABLE_TENNIS ...
	TABLE_TENNIS ActivityType = "TABLE_TENNIS"

	//TENNIS ...
	TENNIS ActivityType = "TENNIS"

	//MEDITATION ...
	MEDITATION ActivityType = "MEDITATION"

	//ACTIVITY_TYPE_OTHER ...
	ACTIVITY_TYPE_OTHER ActivityType = "ACTIVITY_TYPE_OTHER"
)

// ActivityRecordQueryParam ...
type ActivityRecordQueryParam struct {
	ConsumerID   string        `json:"consumerID"`
	ActivityType *ActivityType `json:"activityType"`
	Sort         *SortBy       `json:"sort"`
	Limit        *int32        `json:"limit"`
}

// ActivityRecordCreate ...
type ActivityRecordCreate struct {
	HealthRecordCreate
	ActivityType  ActivityType         `json:"activityType"`
	Code          *CodableConceptInput `json:"code"`
	Frequency     *int32               `json:"frequency"`
	FrequencyUnit *string              `json:"frequencyUnit"`
	Duration      *int32               `json:"duration"`
	DurationUnit  *string              `json:"durationUnit"`
	Distance      *int32               `json:"distance"`
	DistanceUnit  *string              `json:"distanceUnit"`
	Steps         *int32               `json:"steps"`
	Calories      *float64             `json:"calories"`
	CaloryUnit    *string              `json:"caloryUnit"`
	Vigorous      *int32               `json:"vigorous"`
	Moderate      *int32               `json:"moderate"`
	Light         *int32               `json:"light"`
	Sedentary     *int32               `json:"sedentary"`
}

// ActivityRecord ...
type ActivityRecord struct {
	HealthRecord
	Id            string          `json:"id" bson:"_id"`
	ActivityType  ActivityType    `json:"activityType" bson:"activityType"`
	Code          *CodableConcept `json:"code" bson:"code"`
	Frequency     *int32          `json:"frequency" bson:"frequency"`
	FrequencyUnit *string         `json:"frequencyUnit" bson:"frequencyUnit"`
	Duration      *int32          `json:"duration" bson:"duration"`
	DurationUnit  *string         `json:"durationUnit" bson:"durationUnit"`
	Distance      *int32          `json:"distance" bson:"distance"`
	DistanceUnit  *string         `json:"distanceUnit" bson:"distanceUnit"`
	Steps         *int32          `json:"steps" bson:"steps"`
	Calories      *float64        `json:"calories" bson:"calories"`
	CaloryUnit    *string         `json:"caloryUnit" bson:"caloryUnit"`
	Vigorous      *int32          `json:"vigorous" bson:"vigorous"`
	Moderate      *int32          `json:"moderate" bson:"moderate"`
	Light         *int32          `json:"light" bson:"light"`
	Sedentary     *int32          `json:"sedentary" bson:"sedentary"`
	Meta          *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
