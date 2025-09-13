package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// MedicationRecord ..
func (r *Resolver) MedicationRecord(ctx context.Context, args struct {
	ID string
}) (*MedicationRecordResolver, error) {
	medicationRecord, err := ctx.Value(constant.MedicationRecordService).(*service.MedicationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved medication by medication record _ id")

	healthRecordResolver := HealthRecordResolver{&medicationRecord.HealthRecord}
	return &MedicationRecordResolver{healthRecordResolver, medicationRecord}, nil
}

// MedicationRecords ..
func (r *Resolver) MedicationRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*MedicationRecordResolver, error) {
	medicationRecords, err := ctx.Value(constant.MedicationRecordService).(*service.MedicationRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	mrrArr := []*MedicationRecordResolver{}

	for _, mr := range *medicationRecords {
		healthRecordResolver := HealthRecordResolver{&mr.HealthRecord}
		mrr := MedicationRecordResolver{healthRecordResolver, mr}
		mrrArr = append(mrrArr, &mrr)
	}

	return &mrrArr, nil
}
