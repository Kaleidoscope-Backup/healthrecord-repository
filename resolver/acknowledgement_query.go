package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//Acknowledgement ...
func (r *Resolver) Acknowledgement(ctx context.Context, args struct {
	ID string
}) (*AcknowledgementResolver, error) {
	acknowledgement, err := ctx.Value(constant.AcknowledgementService).(*service.AcknowledgementService).FindByID(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &AcknowledgementResolver{acknowledgement}, nil
}
