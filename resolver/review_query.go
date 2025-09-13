package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Review ...
func (r *Resolver) Review(ctx context.Context, args struct {
	ID string
}) (*ReviewResolver, error) {
	review, err := ctx.Value(constant.ReviewService).(*service.ReviewService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &ReviewResolver{review}, nil
}
