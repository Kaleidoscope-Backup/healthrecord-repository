package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//GoalRecord ...
func (r *Resolver) GoalRecord(ctx context.Context, args struct {
	ID string
}) (*GoalRecordResolver, error) {
	goalRecord, err := ctx.Value(constant.GoalRecordService).(*service.GoalRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&goalRecord.HealthRecord}
	return &GoalRecordResolver{healthRecordResolver, goalRecord}, nil
}
