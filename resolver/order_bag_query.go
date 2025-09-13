package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// OrderBag Query
func (r *Resolver) OrderBag(ctx context.Context, args struct {
	ID string
}) (*OrderBagResolver, error) {
	orderBag, err := ctx.Value(constant.OrderBagService).(*service.OrderBagService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved order bag by id : %v", *orderBag)
	return &OrderBagResolver{orderBag}, nil
}
