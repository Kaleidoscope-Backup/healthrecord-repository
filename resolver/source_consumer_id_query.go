package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//SourceConsumerID Query
func (r *Resolver) SourceConsumerID(ctx context.Context, args struct {
	ID string
}) (*SourceConsumerIDResolver, error) {
	sourceConsumerID, err := ctx.Value(constant.SourceConsumerIDService).(*service.SourceConsumerIDService).FindById(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved source record id by _id : %v", *sourceConsumerID)

	return &SourceConsumerIDResolver{sourceConsumerID}, nil
}
