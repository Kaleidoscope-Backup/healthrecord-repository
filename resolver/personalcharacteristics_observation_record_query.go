package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// PersonalCharacteristicsObservationRecord ...
func (r *Resolver) PersonalCharacteristicsObservationRecord(ctx context.Context, args struct {
	ID string
}) (*PersonalCharacteristicsObservationRecordResolver, error) {
	personalCharacteristicsObservationRecord, err := ctx.Value(constant.PersonalCharacteristicsObservationRecordService).(*service.PersonalCharacteristicsObservationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&personalCharacteristicsObservationRecord.HealthRecord}
	return &PersonalCharacteristicsObservationRecordResolver{healthRecordResolver, personalCharacteristicsObservationRecord}, nil
}

// PersonalCharacteristicsObservationRecords ..
func (r *Resolver) PersonalCharacteristicsObservationRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*PersonalCharacteristicsObservationRecordResolver, error) {
	personalCharacteristicsObservationRecords, err := ctx.Value(constant.PersonalCharacteristicsObservationRecordService).(*service.PersonalCharacteristicsObservationRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	pcorrArr := []*PersonalCharacteristicsObservationRecordResolver{}

	for _, pcor := range *personalCharacteristicsObservationRecords {
		healthRecordResolver := HealthRecordResolver{&pcor.HealthRecord}
		pcorr := PersonalCharacteristicsObservationRecordResolver{healthRecordResolver, pcor}
		pcorrArr = append(pcorrArr, &pcorr)
	}

	return &pcorrArr, nil
}
