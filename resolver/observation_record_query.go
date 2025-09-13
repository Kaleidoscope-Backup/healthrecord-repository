package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ObservationRecord Query
func (r *Resolver) ObservationRecord(ctx context.Context, args struct {
	ID string
}) (*ObservationRecordResolver, error) {
	observation, err := ctx.Value(constant.ObservationRecordService).(*service.ObservationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved observation by ID : %v", *observation)

	healthRecordResolver := HealthRecordResolver{&observation.HealthRecord}
	return &ObservationRecordResolver{healthRecordResolver, observation}, nil
}

// ObservationRecords ..
func (r *Resolver) ObservationRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*ObservationRecordResolver, error) {
	observationRecords, err := ctx.Value(constant.ObservationRecordService).(*service.ObservationRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	orrArr := []*ObservationRecordResolver{}

	for _, or := range *observationRecords {
		healthRecordResolver := HealthRecordResolver{&or.HealthRecord}
		orr := ObservationRecordResolver{healthRecordResolver, or}
		orrArr = append(orrArr, &orr)
	}

	return &orrArr, nil
}
