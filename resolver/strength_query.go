package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Strength Query
func (r *Resolver) Strength(ctx context.Context, args struct {
	ID string
}) (*StrengthResolver, error) {
	strength, err := ctx.Value(constant.StrengthService).(*service.StrengthService).FindById(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved strength by strength_id : %v", *strength)
	return &StrengthResolver{strength}, nil
}
