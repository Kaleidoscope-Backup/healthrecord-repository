package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ActivityRecord ...
func (r *Resolver) ActivityRecord(ctx context.Context, args struct {
	ID string
}) (*ActivityRecordResolver, error) {
	activityRecord, err := ctx.Value(constant.ActivityRecordService).(*service.ActivityRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&activityRecord.HealthRecord}
	return &ActivityRecordResolver{healthRecordResolver, activityRecord}, nil
}

// ActivityRecords ..
func (r *Resolver) ActivityRecords(ctx context.Context, args struct {
	Params *model.ActivityRecordQueryParam
}) (*[]*ActivityRecordResolver, error) {
	activityRecords, err := ctx.Value(constant.ActivityRecordService).(*service.ActivityRecordService).FindByParams(args.Params)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	acrrArr := []*ActivityRecordResolver{}

	for _, acr := range *activityRecords {
		healthRecordResolver := HealthRecordResolver{&acr.HealthRecord}
		acrr := ActivityRecordResolver{healthRecordResolver, acr}
		acrrArr = append(acrrArr, &acrr)
	}

	return &acrrArr, nil
}

// ActivityRecordsByConsumerID ..
func (r *Resolver) ActivityRecordsByConsumerID(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*ActivityRecordResolver, error) {
	activityRecords, err := ctx.Value(constant.ActivityRecordService).(*service.ActivityRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	acrrArr := []*ActivityRecordResolver{}

	for _, acr := range *activityRecords {
		healthRecordResolver := HealthRecordResolver{&acr.HealthRecord}
		acrr := ActivityRecordResolver{healthRecordResolver, acr}
		acrrArr = append(acrrArr, &acrr)
	}

	return &acrrArr, nil
}
