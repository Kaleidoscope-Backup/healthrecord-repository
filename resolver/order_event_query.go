package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// OrderEvent Query
func (r *Resolver) OrderEvent(ctx context.Context, args struct {
	ID string
}) (*OrderEventResolver, error) {
	orderEvent, err := ctx.Value(constant.OrderEventService).(*service.OrderEventService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved order by id : %v", *orderEvent)
	return &OrderEventResolver{orderEvent}, nil
}

// OrderEventsByOrder ...
func (r *Resolver) OrderEventsByOrder(ctx context.Context, args struct {
	OrderID string
}) *[]*OrderEventResolver {
	var l []*OrderEventResolver

	//order event records
	orderEventArr, err := ctx.Value(constant.OrderEventService).(*service.OrderEventService).FindEvents(args.OrderID)
	for _, oer := range *orderEventArr {
		oreResolver := OrderEventResolver{oer}
		l = append(l, &oreResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}

// OrderEvents ...
func (r *Resolver) OrderEvents(ctx context.Context, args struct {
	Param *model.OrderEventQueryParam
}) *[]*OrderEventResolver {
	var l []*OrderEventResolver

	//order event records
	orderEventArr, err := ctx.Value(constant.OrderEventService).(*service.OrderEventService).FindByParam(args.Param)
	for _, oer := range *orderEventArr {
		oreResolver := OrderEventResolver{oer}
		l = append(l, &oreResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}
