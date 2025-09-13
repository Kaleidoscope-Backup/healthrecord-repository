package resolver

import (
	logging "github.com/op/go-logging"
	c "gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//CreateSleepRecord ...
func (r *Resolver) CreateSleepRecord(ctx context.Context, args *struct {
	SleepRecord *model.SleepRecordCreate
}) (*SleepRecordResolver, error) {

	sleepRecord := &model.SleepRecord{}
	healthRecord, er := CreateHealthRecord(ctx, &args.SleepRecord.HealthRecordCreate, model.SLEEP)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	sleepRecord.HealthRecord = *healthRecord

	//sleep record properties
	sleepRecord.StartTime = args.SleepRecord.StartTime
	sleepRecord.EndTime = args.SleepRecord.EndTime
	sleepRecord.MainSleep = args.SleepRecord.MainSleep
	sleepRecord.TimeUnit = args.SleepRecord.TimeUnit
	sleepRecord.TotalRecordingTime = args.SleepRecord.TotalRecordingTime
	sleepRecord.TotalSleepTime = args.SleepRecord.TotalSleepTime
	sleepRecord.TimeAwake = args.SleepRecord.TimeAwake
	sleepRecord.SleepEfficiency = args.SleepRecord.SleepEfficiency
	sleepRecord.TimeToFallAsleep = args.SleepRecord.TimeToFallAsleep
	sleepRecord.NumberOfAwekenings = args.SleepRecord.NumberOfAwekenings
	sleepRecord.TimeAfterWakeup = args.SleepRecord.TimeAfterWakeup
	sleepRecord.TimeInBed = args.SleepRecord.TimeInBed

	//sleep stage
	var sleepStageSummary []model.SleepStageInput
	var sleepStages []model.SleepStage
	sleepStageSummary = *args.SleepRecord.SleepStageSummary

	for i := 0; i < len(sleepStageSummary); i++ {
		var sleepStageInput model.SleepStageInput
		sleepStageInput = sleepStageSummary[i]

		slstg := &model.SleepStage{}
		slstg.Type = sleepStageInput.Type
		slstg.Duration = sleepStageInput.Duration
		slstg.TotalSleepTime = sleepStageInput.TotalSleepTime
		slstg.SleepPeriodTime = sleepStageInput.SleepPeriodTime

		sleepStages = append(sleepStages, *slstg)
	}

	sleepRecord.SleepStageSummary = &sleepStages

	//SleepContinuity
	var sleepContinuityInputs []model.SleepContinuityInput
	var sleepContinuities []model.SleepContinuity
	sleepContinuityInputs = *args.SleepRecord.SleepContinuities

	for i := 0; i < len(sleepContinuityInputs); i++ {
		var sleepContinuityInput model.SleepContinuityInput
		sleepContinuityInput = sleepContinuityInputs[i]

		slssct := &model.SleepContinuity{}
		slssct.SourceOfArousal = sleepContinuityInput.SourceOfArousal
		slssct.NermCount = sleepContinuityInput.NermCount
		slssct.NermIndex = sleepContinuityInput.NermIndex
		slssct.RemCount = sleepContinuityInput.RemCount
		slssct.RemIndex = sleepContinuityInput.RemIndex
		slssct.TotalCount = sleepContinuityInput.TotalCount
		slssct.TotalIndex = sleepContinuityInput.TotalIndex

		sleepContinuities = append(sleepContinuities, *slssct)
	}

	sleepRecord.SleepContinuities = &sleepContinuities

	sleepRecord, err := ctx.Value(c.SleepRecordService).(*service.SleepRecordService).CreateSleepRecord(sleepRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created sleep record : %v", *sleepRecord)
	healthRecordResolver := HealthRecordResolver{&sleepRecord.HealthRecord}
	return &SleepRecordResolver{healthRecordResolver, sleepRecord}, nil
}
