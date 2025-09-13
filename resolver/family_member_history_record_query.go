package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// FamilyMemberHistoryRecord ...
func (r *Resolver) FamilyMemberHistoryRecord(ctx context.Context, args struct {
	ID string
}) (*FamilyMemberHistoryRecordResolver, error) {
	familyMemberHistory, err := ctx.Value(constant.FamilyMemberHistoryRecordService).(*service.FamilyMemberHistoryRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved Heart Rate by heartRate_value : %v", *familyMemberHistory)

	healthRecordResolver := HealthRecordResolver{&familyMemberHistory.HealthRecord}
	return &FamilyMemberHistoryRecordResolver{healthRecordResolver, familyMemberHistory}, nil
}
