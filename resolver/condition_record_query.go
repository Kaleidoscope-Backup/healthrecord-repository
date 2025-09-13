package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//ConditionRecord ...
func (r *Resolver) ConditionRecord(ctx context.Context, args struct {
	ID string
}) (*ConditionRecordResolver, error) {
	conditionRecord, err := ctx.Value(constant.ConditionRecordService).(*service.ConditionRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&conditionRecord.HealthRecord}
	return &ConditionRecordResolver{healthRecordResolver, conditionRecord}, nil
}

//ConditionRecords ..
func (r *Resolver) ConditionRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*ConditionRecordResolver, error) {
	conditionRecords, err := ctx.Value(constant.ConditionRecordService).(*service.ConditionRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	crrArr := []*ConditionRecordResolver{}

	for _, cr := range *conditionRecords {
		healthRecordResolver := HealthRecordResolver{&cr.HealthRecord}
		crr := ConditionRecordResolver{healthRecordResolver, cr}
		crrArr = append(crrArr, &crr)
	}

	return &crrArr, nil
}
