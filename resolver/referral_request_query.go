package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//ReferralRequest ...
func (r *Resolver) ReferralRequest(ctx context.Context, args struct {
	ID string
}) (*ReferralRequestResolver, error) {
	refRequest, err := ctx.Value(constant.ReferralRequestService).(*service.ReferralRequestService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &ReferralRequestResolver{refRequest}, nil
}
