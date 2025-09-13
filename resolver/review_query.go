package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Review ...
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
