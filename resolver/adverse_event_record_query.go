package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// AdverseEventRecord ...
func (r *Resolver) AdverseEventRecord(ctx context.Context, args struct {
	ID string
}) (*AdverseEventRecordResolver, error) {
	adverseEventRecord, err := ctx.Value(constant.AdverseEventRecordService).(*service.AdverseEventRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	healthRecordResolver := HealthRecordResolver{&adverseEventRecord.HealthRecord}
	return &AdverseEventRecordResolver{healthRecordResolver, adverseEventRecord}, nil
}
