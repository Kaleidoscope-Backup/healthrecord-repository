package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// SocialHistoryObservationRecord ...
func (r *Resolver) SocialHistoryObservationRecord(ctx context.Context, args struct {
	ID string
}) (*SocialHistoryObservationRecordResolver, error) {
	addictionRecord, err := ctx.Value(constant.SocialHistoryObservationRecordService).(*service.SocialHistoryObservationRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved addiction record by addictionRecord_value : %v", *addictionRecord)

	healthRecordResolver := HealthRecordResolver{&addictionRecord.HealthRecord}
	return &SocialHistoryObservationRecordResolver{healthRecordResolver, addictionRecord}, nil
}

// SocialHistoryObservationRecords ..
func (r *Resolver) SocialHistoryObservationRecords(ctx context.Context, args struct {
	Params *model.SocialHistoryObservationRecordQueryParam
}) (*[]*SocialHistoryObservationRecordResolver, error) {
	socialHistoryObservationRecords, err := ctx.Value(constant.SocialHistoryObservationRecordService).(*service.SocialHistoryObservationRecordService).FindByParams(args.Params)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	shorrArr := []*SocialHistoryObservationRecordResolver{}

	for _, shor := range *socialHistoryObservationRecords {
		healthRecordResolver := HealthRecordResolver{&shor.HealthRecord}
		shorr := SocialHistoryObservationRecordResolver{healthRecordResolver, shor}
		shorrArr = append(shorrArr, &shorr)
	}

	return &shorrArr, nil
}

// SocialHistoryObservationRecordsByConsumerID ..
func (r *Resolver) SocialHistoryObservationRecordsByConsumerID(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*SocialHistoryObservationRecordResolver, error) {
	socialHistoryObservationRecords, err := ctx.Value(constant.SocialHistoryObservationRecordService).(*service.SocialHistoryObservationRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	shorrArr := []*SocialHistoryObservationRecordResolver{}

	for _, shor := range *socialHistoryObservationRecords {
		healthRecordResolver := HealthRecordResolver{&shor.HealthRecord}
		shorr := SocialHistoryObservationRecordResolver{healthRecordResolver, shor}
		shorrArr = append(shorrArr, &shorr)
	}

	return &shorrArr, nil
}
