package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// HealthRecord ...
func (r *Resolver) HealthRecord(ctx context.Context, args struct {
	ID string
}) (*HealthRecordResolver, error) {
	healthRecord, err := ctx.Value(constant.HealthRecordService).(*service.HealthRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved Heart Rate by health record : %v", *healthRecord)

	healthRecordResolver := HealthRecordResolver{healthRecord}
	return &healthRecordResolver, nil
}

// HealthRecords ..
func (r *Resolver) HealthRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*HealthRecordResolver, error) {
	healthRecords, err := ctx.Value(constant.HealthRecordService).(*service.HealthRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	hrrArr := []*HealthRecordResolver{}

	for _, hr := range *healthRecords {
		healthRecordResolver := HealthRecordResolver{hr}
		hrrArr = append(hrrArr, &healthRecordResolver)
	}

	return &hrrArr, nil
}
