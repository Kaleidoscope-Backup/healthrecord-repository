package resolver

import (
	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateSourceOrganizationID ..
func (r *Resolver) CreateSourceOrganizationID(ctx context.Context, args *struct {
	SourceID string
	Type     string
}) (*SourceOrganizationIDResolver, error) {
	sourceOrganizationID := &model.SourceOrganizationID{}
	sourceOrganizationID.SourceID = args.SourceID
	sourceOrganizationID.Type = &args.Type

	sourceOrganizationID, err := ctx.Value(constant.SourceOrganizationIDService).(*service.SourceOrganizationIDService).
		CreateSourceOrganizationID(sourceOrganizationID)

	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created source record id : %v", *sourceOrganizationID)
	return &SourceOrganizationIDResolver{sourceOrganizationID}, nil
}
