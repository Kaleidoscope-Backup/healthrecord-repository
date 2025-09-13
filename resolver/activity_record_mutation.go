package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateActivityRecord creates a new ActivityRecord in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateActivityRecord(ctx context.Context, args *struct {
	ActivityRecord *model.ActivityRecordCreate
}) (*ActivityRecordResolver, error) {

	activityRecord := &model.ActivityRecord{}

	//populate activity record object
	activityRecord.ActivityType = args.ActivityRecord.ActivityType
	activityRecord.Duration = args.ActivityRecord.Duration
	activityRecord.DurationUnit = args.ActivityRecord.DurationUnit
	activityRecord.Distance = args.ActivityRecord.Distance
	activityRecord.DistanceUnit = args.ActivityRecord.DistanceUnit
	activityRecord.Steps = args.ActivityRecord.Steps
	activityRecord.Calories = args.ActivityRecord.Calories
	activityRecord.CaloryUnit = args.ActivityRecord.CaloryUnit
	activityRecord.Vigorous = args.ActivityRecord.Vigorous
	activityRecord.Moderate = args.ActivityRecord.Moderate
	activityRecord.Light = args.ActivityRecord.Light
	activityRecord.Sedentary = args.ActivityRecord.Sedentary
	activityRecord.Frequency = args.ActivityRecord.Frequency
	activityRecord.FrequencyUnit = args.ActivityRecord.FrequencyUnit

	if args.ActivityRecord.Code != nil {
		code, er := CreateCodableConceptFromInput(ctx, args.ActivityRecord.Code)
		if er != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
			return nil, er
		}
		activityRecord.Code = code
	}

	var meta models.Meta
	meta.VersionId = "0.0.1"
	activityRecord.Meta = &meta

	healthRecord, er := CreateHealthRecord(ctx, &args.ActivityRecord.HealthRecordCreate, model.ACTIVITY)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	activityRecord.HealthRecord = *healthRecord
	activityRecord, err := ctx.Value(constant.ActivityRecordService).(*service.ActivityRecordService).CreateActivityRecord(activityRecord)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created activity Record : %v", *activityRecord)

	healthRecordResolver := HealthRecordResolver{&activityRecord.HealthRecord}
	return &ActivityRecordResolver{healthRecordResolver, activityRecord}, nil
}
