package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// AllergyRecord ...
func (r *Resolver) AllergyRecord(ctx context.Context, args struct {
	ID string
}) (*AllergyRecordResolver, error) {
	allergyRecord, err := ctx.Value(constant.AllergyRecordService).(*service.AllergyRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&allergyRecord.HealthRecord}
	return &AllergyRecordResolver{healthRecordResolver, allergyRecord}, nil
}

// AllergyRecords ..
func (r *Resolver) AllergyRecords(ctx context.Context, args struct {
	ConsumerID string
}) (*[]*AllergyRecordResolver, error) {
	allergyRecords, err := ctx.Value(constant.AllergyRecordService).(*service.AllergyRecordService).FindByConsumerID(args.ConsumerID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	arrArr := []*AllergyRecordResolver{}

	for _, ar := range *allergyRecords {
		healthRecordResolver := HealthRecordResolver{&ar.HealthRecord}
		arr := AllergyRecordResolver{healthRecordResolver, ar}
		arrArr = append(arrArr, &arr)
	}

	return &arrArr, nil
}
