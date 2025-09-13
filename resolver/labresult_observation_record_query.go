package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// LabResultObservationRecord ...
func (r *Resolver) LabResultObservationRecord(ctx context.Context, args struct {
	ID string
}) (*LabResultObservationRecordResolver, error) {
	labResultObservationRecord, err := ctx.Value(constant.LabResultObservationRecordService).(*service.LabResultObservationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved lab result observations : %v", *labResultObservationRecord)

	healthRecordResolver := HealthRecordResolver{&labResultObservationRecord.HealthRecord}
	return &LabResultObservationRecordResolver{healthRecordResolver, labResultObservationRecord}, nil
}

// LabResultObservationRecords ..
func (r *Resolver) LabResultObservationRecords(ctx context.Context, args struct {
	Param *model.LabResultObservationRecordQueryParam
}) (*[]*LabResultObservationRecordResolver, error) {
	labResultObservationRecords, err := ctx.Value(constant.LabResultObservationRecordService).(*service.LabResultObservationRecordService).FindByParam(args.Param)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	lorrArr := []*LabResultObservationRecordResolver{}

	for _, lor := range *labResultObservationRecords {
		healthRecordResolver := HealthRecordResolver{&lor.HealthRecord}
		lorr := LabResultObservationRecordResolver{healthRecordResolver, lor}
		lorrArr = append(lorrArr, &lorr)
	}

	return &lorrArr, nil
}

// LabResultObservationRecordsByConsumerID ..
func (r *Resolver) LabResultObservationRecordsByConsumerID(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*LabResultObservationRecordResolver, error) {
	labResultObservationRecords, err := ctx.Value(constant.LabResultObservationRecordService).(*service.LabResultObservationRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	lorrArr := []*LabResultObservationRecordResolver{}

	for _, lor := range *labResultObservationRecords {
		healthRecordResolver := HealthRecordResolver{&lor.HealthRecord}
		lorr := LabResultObservationRecordResolver{healthRecordResolver, lor}
		lorrArr = append(lorrArr, &lorr)
	}

	return &lorrArr, nil
}
