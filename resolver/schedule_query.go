package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Schedule ...
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
