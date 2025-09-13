package resolver

import (
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

//SourceOrganizationID Query
func (r *Resolver) SourceOrganizationID(ctx context.Context, args struct {
	ID string
}) (*SourceOrganizationIDResolver, error) {
	sourceOrganizationID, err := ctx.Value(constant.SourceOrganizationIDService).(*service.SourceOrganizationIDService).FindById(args.ID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved source record id by _id : %v", *sourceOrganizationID)

	return &SourceOrganizationIDResolver{sourceOrganizationID}, nil
}
