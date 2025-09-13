package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Procedure Query
func (r *Resolver) ProcedureRecord(ctx context.Context, args struct {
	ID string
}) (*ProcedureRecordResolver, error) {
	practitioner, err := ctx.Value(constant.ProcedureRecordService).(*service.ProcedureRecordService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved procedure by procedure_id : %v", *practitioner)

	healthRecordResolver := HealthRecordResolver{&practitioner.HealthRecord}
	return &ProcedureRecordResolver{healthRecordResolver, practitioner}, nil
}
