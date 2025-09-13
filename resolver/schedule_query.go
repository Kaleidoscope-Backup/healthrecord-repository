package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Schedule ...
func (r *Resolver) Schedule(ctx context.Context, args struct {
	ID string
}) (*ScheduleResolver, error) {
	schedule, err := ctx.Value(constant.ScheduleService).(*service.ScheduleService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &ScheduleResolver{schedule}, nil
}
