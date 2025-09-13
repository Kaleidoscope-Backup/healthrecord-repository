package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// ReferralRequest ...
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
