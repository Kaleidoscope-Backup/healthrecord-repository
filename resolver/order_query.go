package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Order Query
func (r *Resolver) Order(ctx context.Context, args struct {
	ID string
}) (*OrderResolver, error) {
	order, err := ctx.Value(constant.OrderService).(*service.OrderService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved order by id : %v", *order)
	return &OrderResolver{order}, nil
}

// Orders ...
func (r *Resolver) Orders(ctx context.Context, args struct {
	Param *model.OrderQueryParam
}) *[]*OrderResolver {
	var l []*OrderResolver

	orderArr, err := ctx.Value(constant.OrderService).(*service.OrderService).FindOrders(args.Param)
	for _, or := range *orderArr {
		orResolver := OrderResolver{or}
		l = append(l, &orResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}

// OrdersBySupplier ...
func (r *Resolver) OrdersBySupplier(ctx context.Context, args struct {
	SupplierID string
}) *[]*OrderResolver {
	var l []*OrderResolver

	orderArr, err := ctx.Value(constant.OrderService).(*service.OrderService).FindOrdersBySupplier(args.SupplierID)
	for _, or := range *orderArr {
		orResolver := OrderResolver{or}
		l = append(l, &orResolver)
	}

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil
	}

	return &l
}
