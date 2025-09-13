package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//ImmunizationRecord ...
func (r *Resolver) ImmunizationRecord(ctx context.Context, args struct {
	ID string
}) (*ImmunizationRecordResolver, error) {
	immunizationRecord, err := ctx.Value(constant.ImmunizationRecordService).(*service.ImmunizationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved immunization: %v", *immunizationRecord)

	healthRecordResolver := HealthRecordResolver{&immunizationRecord.HealthRecord}
	return &ImmunizationRecordResolver{healthRecordResolver, immunizationRecord}, nil
}

//ImmunizationRecords ..
func (r *Resolver) ImmunizationRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*ImmunizationRecordResolver, error) {
	immunizationRecords, err := ctx.Value(constant.ImmunizationRecordService).(*service.ImmunizationRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	irrArr := []*ImmunizationRecordResolver{}

	for _, ir := range *immunizationRecords {
		healthRecordResolver := HealthRecordResolver{&ir.HealthRecord}
		irr := ImmunizationRecordResolver{healthRecordResolver, ir}
		irrArr = append(irrArr, &irr)
	}

	return &irrArr, nil
}
