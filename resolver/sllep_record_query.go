package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// SleepRecord ...
func (r *Resolver) SleepRecord(ctx context.Context, args struct {
	ID string
}) (*SleepRecordResolver, error) {
	sleepRecord, err := ctx.Value(constant.SleepRecordService).(*service.SleepRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved Vital by vital_id : %v", *sleepRecord)
	healthRecordResolver := HealthRecordResolver{&sleepRecord.HealthRecord}
	return &SleepRecordResolver{healthRecordResolver, sleepRecord}, nil
}
